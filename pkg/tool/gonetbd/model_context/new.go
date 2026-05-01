package model_context

import (
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/constant"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *netbox.Client,
	h *sentry.Hub,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		client: c,
		hub:    h,
	}
	result.register()

	return result
}
