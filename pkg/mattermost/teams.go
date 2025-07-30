package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Teams(userIdentifier string) []*model.Team {
	result, r, e := c.client.GetTeamsForUser(
		c.context,
		userIdentifier,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
