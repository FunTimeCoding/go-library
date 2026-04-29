package model_context

import (
	"github.com/funtimecoding/go-library/pkg/firefox"
	"github.com/getsentry/sentry-go"
)

func New(
	c *firefox.Client,
	h *sentry.Hub,
) *Server {
	return &Server{client: c, hub: h}
}
