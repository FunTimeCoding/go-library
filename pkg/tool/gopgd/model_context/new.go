package model_context

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			constant.Version,
			server.WithToolCapabilities(true),
		),
		store:    s,
		reporter: r,
	}
	result.register()

	return result
}
