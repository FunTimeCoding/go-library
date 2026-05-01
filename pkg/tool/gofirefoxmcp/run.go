package gofirefoxmcp

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Firefox,
	h *sentry.Hub,
) {
	c := firefox.NewEnvironment()
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.BridgePort),
			func(m *http.ServeMux) {
				m.Handle("/", c)
			},
			web.RecoveryMiddleware(h),
		),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(c, h).Mount(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
