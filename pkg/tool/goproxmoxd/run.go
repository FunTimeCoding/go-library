package goproxmoxd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
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
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				t := telemetry.NewEnvironment()
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(v, r),
						[]generated.StrictMiddlewareFunc{
							web.TelemetryMiddleware(t),
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
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
