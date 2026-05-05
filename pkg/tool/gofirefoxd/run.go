package gofirefoxd

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func Run(
	o *option.Firefox,
	r face.Reporter,
) {
	c := firefox.NewEnvironment()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.BridgePort),
			func(m *http.ServeMux) {
				m.Handle("/", c)
			},
			web.RecoveryMiddleware(r),
		),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(c, r, o.Version).Mount(m)
			},
			web.RecoveryMiddleware(r),
		),
	).RunUntilSignal()
}
