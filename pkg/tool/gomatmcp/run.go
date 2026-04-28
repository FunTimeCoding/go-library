package gomatmcp

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/lifecycle"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/tool"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"net/http"
)

func Run(
	o *option.Mattermost,
	h *sentry.Hub,
) {
	l := logger.New(context.Background())
	c := mattermost.NewEnvironment()
	var m *monitor.Monitor
	var p []lifecycle.Option

	if v := monitor.LoadConfiguration(); v.Enabled {
		m = monitor.New(c, v, l, h)
		p = append(p, lifecycle.WithWorker(m))
	}

	lifecycle.New(
		l,
		append(
			p,
			lifecycle.WithServer(
				web.AddressPort(o.Port),
				func(u *http.ServeMux) {
					s := server.NewMCPServer(
						"Mattermost MCP",
						constant.DefaultVersion,
					)
					addTool(s, tool.New(c, m), m != nil)
					generative.New(s).Setup(u)
				},
			),
		)...,
	).RunUntilSignal()
}
