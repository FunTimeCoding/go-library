package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) TeamChannel(name string) *model.Channel {
	return c.ChannelByName(c.team, name)
}
