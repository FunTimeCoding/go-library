package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	p *poller.Poller,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			"alert-log",
			constant.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		store:  s,
		poller: p,
	}
	result.register()

	return result
}
