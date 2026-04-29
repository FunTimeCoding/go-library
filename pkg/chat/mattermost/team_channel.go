package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) TeamChannel(name string) (*model.Channel, error) {
	return c.ChannelByName(c.team, name)
}
