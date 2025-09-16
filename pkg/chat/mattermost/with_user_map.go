package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/user_map"

func WithUserMap(m *user_map.Map) Option {
	return func(c *Client) {
		c.user = m
	}
}
