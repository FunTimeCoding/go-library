package model_context

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	c *sentry.Client,
	organization string,
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
		client:       c,
		organization: organization,
		reporter:     r,
	}
	result.register()

	return result
}
