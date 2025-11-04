package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Team(name string) *model.Team {
	result, r, e := c.client.GetTeamByName(
		c.context,
		name,
		constant.EmptyEntityTag,
	)
	panicOnError(r, e)

	return result
}
