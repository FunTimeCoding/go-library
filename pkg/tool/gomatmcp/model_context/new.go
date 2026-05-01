package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	h *sentry.Hub,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
		),
		client:  m,
		monitor: o,
		hub:     h,
	}
	result.register()

	return result
}
