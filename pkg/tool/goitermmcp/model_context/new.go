package model_context

import (
	"github.com/funtimecoding/go-library/pkg/iterm"
	"github.com/getsentry/sentry-go"
)

func New(
	c *iterm.Client,
	h *sentry.Hub,
) *Server {
	return &Server{client: c, hub: h}
}
