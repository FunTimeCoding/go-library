package goclauded

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/event/notifier"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	generated "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/option"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/session_indexer"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/watcher"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/web"
	memory "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/connect"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/connect"
	library "github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func Run(
	o *option.Option,
	r face.Reporter,
) {
	start := time.Now()
	elapsed := func() float64 { return time.Since(start).Seconds() }
	l := logger.New(context.Background())
	n := notifier.New()
	s := store.New(store.DefaultDatabasePath(), time.Now)
	l.Structured("store_ready", "elapsed", elapsed())
	h := claude.New().Base()
	result := sweep.Run(h)
	l.Structured(
		"sweep_complete",
		"elapsed",
		elapsed(),
		"copied",
		result.Copied,
		"updated",
		result.Updated,
		"skipped",
		result.Skipped,
	)
	s.ClearBindings()
	memoryClient := memory.Wait(l)
	l.Structured("gomemoryd_connected", "elapsed", elapsed())
	queryClient := connect.Wait(l)
	l.Structured("goqueryd_connected", "elapsed", elapsed())
	summaryIdx := session_indexer.New(
		queryClient,
		"summaries",
		"session-summary",
	)
	completionIdx := session_indexer.New(
		queryClient,
		"completions",
		"session-completion",
	)
	v := service.New(
		s,
		claude.New(),
		memoryClient,
		summaryIdx,
		completionIdx,
		n,
		r,
		h,
		time.Now,
		l,
	)
	v.ReconcileSummaries()
	l.Structured("summaries_reconciled", "elapsed", elapsed())
	v.ReconcileCompletions()
	l.Structured("completions_reconciled", "elapsed", elapsed())
	v.PopulateCache()
	l.Structured("cache_populated", "elapsed", elapsed())
	v.BackfillSessions()
	l.Structured("sessions_backfilled", "elapsed", elapsed())
	v.CheckConsistency()
	l.Structured("consistency_checked", "elapsed", elapsed())
	go v.RunTimeoutLoop()
	w := watcher.New(v, l, r, h)

	if environment.Exists(constant.MonitorUsageEnvironment) {
		go v.RunUsageLoop()
	}

	lifecycle.New(
		l,
		lifecycle.WithWorker(w),
		lifecycle.WithServerMiddleware(
			library.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(v, l, r, h, o.SessionExportPath),
						nil,
					),
					m,
				)
				model_context.New(
					v,
					r,
					l,
					telemetry.NewEnvironment(),
					o.Version,
				).Mount(m)
				web.New(v).Mount(m)
			},
			library.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
