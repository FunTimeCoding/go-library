package model_context

import (
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/constant"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(c *habitica.Client, h *sentry.Hub) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		habitica: c,
		hub:      h,
	}
	result.register()

	return result
}
