package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/mark3labs/mcp-go/server"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	c *gitlab.Client,
	r face.Reporter,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
		),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
