package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) DefaultChannel() *model.Channel {
	return c.channel
}
