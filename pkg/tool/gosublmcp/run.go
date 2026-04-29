package gosublmcp

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(
	o *option.Sublime,
	h *sentry.Hub,
) {
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				s := server.NewMCPServer(constant.Name, constant.Version)
				addTool(s, model_context.New(sublime.NewEnvironment(), h))
				generative.New(s).Setup(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
