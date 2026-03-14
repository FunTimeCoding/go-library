package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(s *store.Store) *Server {
	result := &Server{
		server: server.NewMCPServer(
			"maintenance-log",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		store: s,
	}
	result.register()

	return result
}
