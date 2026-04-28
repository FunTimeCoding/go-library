package gogitlabmcp

import (
	"context"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/option"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/tool"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(
	o *option.Gitlab,
	h *sentry.Hub,
) {
	s := server.NewMCPServer(constant.Name, constant.Version)
	addTool(s, tool.New(gitlab.NewEnvironment().Nested()))
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(m *http.ServeMux) {
				generative.New(s).Setup(m)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
