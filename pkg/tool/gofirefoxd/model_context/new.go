package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *firefox.Client,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.ServerInstructions),
		),
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
