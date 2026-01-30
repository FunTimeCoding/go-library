package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Channel(identifier string) *model.Channel {
	result, r, e := c.client.GetChannel(c.context, identifier)
	panicOnError(r, e)

	return result
}
