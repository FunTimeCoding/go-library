package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Channels(
	t *model.Team,
	u *model.User,
) []*model.Channel {
	result, r, e := c.client.GetChannelsForTeamForUser(
		t.Id,
		u.Id,
		false,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
