package gopgd

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	generated "github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Postgres,
	h *sentry.Hub,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				s := store.New(o.Inventory)
				generated.HandlerFromMux(server.New(s), m)
				generative.New(model_context.New(s).Nested()).Setup(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
