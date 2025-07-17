package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) ChannelByName(
	t *model.Team,
	name string,
) *model.Channel {
	result, r, e := c.client.GetChannelByName(
		name,
		t.Id,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
