package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *service.Service,
	r face.Reporter,
	l *logger.Logger,
	t face.Recorder,
	version string,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		service:   s,
		reporter:  r,
		logger:    l,
		telemetry: t,
	}
	result.register()
	result.registerResources()

	return result
}
