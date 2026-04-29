package model_context

import (
	"github.com/funtimecoding/go-library/pkg/sublime"
	"github.com/getsentry/sentry-go"
)

func New(
	c *sublime.Client,
	h *sentry.Hub,
) *Server {
	return &Server{client: c, hub: h}
}
