package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
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
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
