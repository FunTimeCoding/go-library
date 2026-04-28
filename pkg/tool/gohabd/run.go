package gohabd

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/route"
	generated "github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Habitica,
	h *sentry.Hub,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				c := habitica.NewEnvironment()
				generated.HandlerFromMux(route.New(c), m)
				generative.New(model_context.New(c).Nested()).Setup(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
