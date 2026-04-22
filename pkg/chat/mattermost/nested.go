package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Nested() *model.Client4 {
	return c.client
}
