package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) AllChannels() []*model.ChannelWithTeamData {
	result, r, e := c.client.GetAllChannels(
		c.context,
		0,
		constant.PerPage,
		constant.EmptyEntityTag,
	)
	panicOnError(r, e)

	return result
}
