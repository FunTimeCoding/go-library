package gohabiticad

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
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
				generated.HandlerFromMux(server.New(c), m)
				model_context.New(c, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
