package gomaintlogd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	maintenanceWeb "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Run(
	o *option.Log,
	r face.Reporter,
) {
	g := logger.New(context.Background())
	s := store.New(o.PostgresLocator, o.LitePath)
	defer s.Close()
	lifecycle.New(
		g,
		lifecycle.WithServer(
			lifecycleServer.New(
				constant.ListenAddress,
				func(m *http.ServeMux) {
					t := telemetry.NewEnvironment()
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, r),
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
						s,
						r,
						t,
						o.Version,
					).Mount(m)
					maintenanceWeb.New(s).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
