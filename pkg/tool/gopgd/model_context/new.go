package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/store"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	h *sentry.Hub,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			constant.Version,
			server.WithToolCapabilities(true),
		),
		store: s,
		hub:   h,
	}
	result.register()

	return result
}
