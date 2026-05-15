package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosublimed/constant"
	sublime "github.com/funtimecoding/go-library/pkg/tool/gosublimed/face"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c sublime.SublimeSource,
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
