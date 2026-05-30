package gohabiticad

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/telemetry"
	generated "github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/option"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Habitica,
	r face.Reporter,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				c := habitica.NewEnvironment()
				t := telemetry.NewEnvironment()
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(c),
						[]generated.StrictMiddlewareFunc{
							web.TelemetryMiddleware(t),
						},
					),
					m,
				)
				model_context.New(c, r, t, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
