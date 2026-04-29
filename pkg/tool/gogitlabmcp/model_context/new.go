package model_context

import (
	"github.com/getsentry/sentry-go"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(
	c *gitlab.Client,
	h *sentry.Hub,
) *Server {
	return &Server{client: c, hub: h}
}
