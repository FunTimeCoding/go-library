package assistant

import "github.com/getsentry/sentry-go"

func WithSentry(h *sentry.Hub) Option {
	return func(c *Client) {
		c.hub = h
	}
}
