package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabd/constant"
	"github.com/mark3labs/mcp-go/server"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	c *gitlab.Client,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			version,
		),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
