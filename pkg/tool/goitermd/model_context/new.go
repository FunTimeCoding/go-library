package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goitermd/constant"
	iterm "github.com/funtimecoding/go-library/pkg/tool/goitermd/face"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c iterm.ItermSource,
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
