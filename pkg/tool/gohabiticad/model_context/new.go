package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *habitica.Client,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		habitica: c,
		reporter: r,
	}
	result.register()

	return result
}
