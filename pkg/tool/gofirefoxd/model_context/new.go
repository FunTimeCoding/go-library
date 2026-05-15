package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/constant"
	firefox "github.com/funtimecoding/go-library/pkg/tool/gofirefoxd/face"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c firefox.FirefoxSource,
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
		client:   c,
		reporter: r,
	}
	result.register()

	return result
}
