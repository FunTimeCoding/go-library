package gogitlabmcp

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
)

func Run(
	o *option.Gitlab,
	h *sentry.Hub,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				model_context.New(
					gitlab.NewEnvironment().Nested(),
					h,
				).Mount(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
