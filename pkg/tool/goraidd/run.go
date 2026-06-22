package goraidd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	lifecycleServer "github.com/funtimecoding/go-library/pkg/lifecycle/server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/raid_parser"
	"github.com/funtimecoding/go-library/pkg/relational"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/option"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	raidWeb "github.com/funtimecoding/go-library/pkg/tool/goraidd/web"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Raid,
	r face.Reporter,
) {
	l := logger.New(context.Background())
	s := store.New(
		relational.New(o.PostgresLocator).Mapper(),
		o.LogCachePath,
		o.ElitePath,
		l,
		r,
	)
	lifecycle.New(
		l,
		lifecycle.WithWorker(s),
		lifecycle.WithServer(
			lifecycleServer.New(
				web.AddressPort(o.Port),
				func(m *http.ServeMux) {
					p := raid_parser.New("localhost:8081", true)
					generated.HandlerFromMux(
						generated.NewStrictHandler(
							server.New(s, o.OutputPath, r),
							[]generated.StrictMiddlewareFunc{
								func(
									f generated.StrictHandlerFunc,
									operation string,
								) generated.StrictHandlerFunc {
									t := telemetry.NewEnvironment()

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
					raidWeb.New(
						s,
						o.ElitePath,
						o.OutputPath,
						p,
					).Mount(m)
				},
			).WithMiddleware(web.RecoveryMiddleware(r)),
		),
	).RunUntilSignal()
}
