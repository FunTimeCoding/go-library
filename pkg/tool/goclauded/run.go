package goclauded

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/notifier"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/option"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/summary_indexer"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
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
	l := logger.New(context.Background())
	n := notifier.New()
	s := store.New(store.DefaultDatabasePath(), time.Now)
	h := claude.New().Base()
	sweep.Run(h)
	s.ClearBindings()
	v := service.New(
		s,
		claude.New(),
		memory.Wait(l),
		summary_indexer.New(connect.Wait(l)),
		n,
		time.Now,
		l,
	)
	v.ReconcileSummaries()
	v.BackfillSessions()
	v.CheckConsistency()
	go v.RunTimeoutLoop()
	go sweep.RunLoop(h)
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			library.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					server.New(v, l, h, o.SessionExportPath),
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
