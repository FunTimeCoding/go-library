package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) ChannelByName(
	t *model.Team,
	name string,
) *model.Channel {
	result, r, e := c.client.GetChannelByName(
		c.context,
		name,
		t.Id,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
