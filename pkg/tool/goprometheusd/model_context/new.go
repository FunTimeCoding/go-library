package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/service"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	v *service.Service,
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
		service:  v,
		reporter: r,
	}
	result.register()

	return result
}
