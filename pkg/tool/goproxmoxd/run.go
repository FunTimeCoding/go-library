package goproxmoxd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Proxmox,
	r face.Reporter,
) {
	v := service.New(o.Inventory)
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			lifecycleServer.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(v, r),
							[]generated.StrictMiddlewareFunc{
								func(
									f generated.StrictHandlerFunc,
									operation string,
								) generated.StrictHandlerFunc {
									return func(
										x context.Context,
										w http.ResponseWriter,
										r *http.Request,
										request any,
									) (any, error) {
										response, e := f(x, w, r, request)
										web.RecordTelemetry(t, operation, e)

										return response, e
									}
								},
							},
						),
						m,
					)
					model_context.New(
						v,
						r,
						t,
						o.Version,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
