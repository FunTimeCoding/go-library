package model_context

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost"
	"github.com/funtimecoding/go-library/pkg/tool/gomatmcp/monitor"
	"github.com/getsentry/sentry-go"
)

func New(
	m *mattermost.Client,
	o *monitor.Monitor,
	h *sentry.Hub,
) *Server {
	return &Server{client: m, monitor: o, hub: h}
}
