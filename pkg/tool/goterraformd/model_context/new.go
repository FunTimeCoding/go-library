package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/runner"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/store"
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
			constant.Identity.Name(),
			version,
			server.WithToolCapabilities(true),
			server.WithInstructions(constant.Identity.Instructions()),
		),
		runner:   n,
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
