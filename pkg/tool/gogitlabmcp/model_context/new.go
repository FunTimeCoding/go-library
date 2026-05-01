package model_context

import (
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	c *gitlab.Client,
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
