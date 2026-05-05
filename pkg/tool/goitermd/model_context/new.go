package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/iterm"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *iterm.Client,
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
