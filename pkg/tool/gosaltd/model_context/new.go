package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/store"
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
		),
		runner:   n,
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
