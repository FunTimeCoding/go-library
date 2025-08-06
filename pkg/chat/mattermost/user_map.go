package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/user_map"

func (c *Client) UserMap() *user_map.Map {
	return c.user
}
