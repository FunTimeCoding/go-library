package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) AllChannels() ([]*model.ChannelWithTeamData, error) {
	result, _, e := c.client.GetAllChannels(
		c.context,
		0,
		constant.PerPage,
		constant.EmptyEntityTag,
	)

	return result, e
}
