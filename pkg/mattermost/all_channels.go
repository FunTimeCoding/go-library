package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) AllChannels() []*model.ChannelWithTeamData {
	result, r, e := c.client.GetAllChannels(
		0,
		100,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
