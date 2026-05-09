package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	n *runner.Runner,
	s *store.Store,
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
		runner:   n,
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
