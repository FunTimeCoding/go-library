package model_context

import (
	modelContext "github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/poller"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/getsentry/sentry-go"
	"github.com/mark3labs/mcp-go/server"
)

func New(
	s *store.Store,
	p *poller.Poller,
	h *sentry.Hub,
) *Server {
	result := &Server{
		server: server.NewMCPServer(
			constant.Name,
			modelContext.DefaultVersion,
			server.WithToolCapabilities(true),
		),
		store:  s,
		poller: p,
		hub:    h,
	}
	result.register()

	return result
}
