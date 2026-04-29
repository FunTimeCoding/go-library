package gofirefoxmcp

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/firefox"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxmcp/option"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(
	o *option.Firefox,
	h *sentry.Hub,
) {
	l := logger.New(context.Background())
	c := firefox.NewEnvironment()
	s := server.NewMCPServer(constant.Name, constant.Version)
	t := model_context.New(c, h)
	addTool(s, t)
	v := generative.New(s)
	lifecycle.New(
		l,
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.BridgePort),
			func(mx *http.ServeMux) {
				mx.Handle("/", c)
			},
			web.RecoveryMiddleware(h),
		),
		lifecycle.WithServerMiddleware(
			web.AddressPort(o.Port),
			func(mx *http.ServeMux) {
				v.Setup(mx)
			},
			web.RecoveryMiddleware(h),
		),
	).RunUntilSignal()
}
