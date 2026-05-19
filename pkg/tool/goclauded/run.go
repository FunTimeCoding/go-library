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
	s := store.New(store.DefaultDatabasePath(), n, time.Now)
	h := claude.New().Base()
	sweep.Run(h)
	c := claude.New()
	y := memory.Wait(l)
	q := connect.Wait(l)
	s.ClearBindings()
	v := service.New(s, y)
	i := summary_indexer.New(q)
	reconcileSummaries(s, i)
	go s.RunCleanup()
	go s.RunTimeoutLoop()
	go sweep.RunLoop(h)
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			library.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					server.New(v, c, l, h, o.SessionExportPath),
					m,
				)
				model_context.New(
					v,
					i,
					r,
					l,
					telemetry.NewEnvironment(),
					o.Version,
				).Mount(m)
				web.New(v, n, c).Mount(m)
			},
			library.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
