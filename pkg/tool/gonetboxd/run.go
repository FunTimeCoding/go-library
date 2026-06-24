package gonetboxd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Netbox,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(o.Client, r),
							[]generated.StrictMiddlewareFunc{
								func(
									f generated.StrictHandlerFunc,
									operation string,
								) generated.StrictHandlerFunc {
									return func(
										x context.Context,
										w http.ResponseWriter,
										q *http.Request,
										request any,
									) (any, error) {
										response, e := f(x, w, q, request)
										web.RecordTelemetry(t, operation, e)

										return response, e
									}
								},
							},
						),
						m,
					)
					model_context.New(
						o.Client,
						r,
						t,
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
