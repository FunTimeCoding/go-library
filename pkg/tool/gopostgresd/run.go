package gopostgresd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Postgres,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					s := store.New(o.Inventory)
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, r),
							[]generated.StrictMiddlewareFunc{
								web.TelemetryMiddleware(t),
							},
						),
						m,
					)
					model_context.New(
						s,
						r,
						t,
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
