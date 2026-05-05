package gonetboxd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
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
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generated.HandlerFromMux(server.New(o.Client), m)
				model_context.New(o.Client, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
