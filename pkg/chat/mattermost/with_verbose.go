package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/verbose_transport"

func WithVerbose(v bool) Option {
	return func(c *Client) {
		if v {
			c.client.HTTPClient = verbose_transport.New()
		}
	}
}
