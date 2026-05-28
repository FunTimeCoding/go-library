package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/service"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *service.Service,
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
		service:  s,
		reporter: r,
	}
	result.register()
	result.registerResources()

	return result
}
