package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	svc *service.Service,
	readOnly bool,
	r face.Reporter,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithResourceCapabilities(true, false),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		service:  svc,
		readOnly: readOnly,
		reporter: r,
	}
	result.register()
	result.registerResources()

	if !readOnly {
		result.registerWrite()
	}

	return result
}
