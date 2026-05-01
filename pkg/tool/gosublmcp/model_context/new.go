package model_context

import (
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/funtimecoding/go-library/pkg/tool/gosublmcp/constant"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *sublime.Client,
	h *sentry.Hub,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
		),
		client: c,
		hub:    h,
	}
	result.register()

	return result
}
