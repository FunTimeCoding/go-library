package mattermost

import "github.com/funtimecoding/go-library/pkg/mattermost/user_map"

func WithUserMap(m *user_map.Map) OptionFunc {
	return func(c *Client) {
		c.user = m
	}
}
