package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Team(name string) *model.Team {
	result, r, e := c.client.GetTeamByName(name, constant.EmptyEntityTag)
	panicOnError(e, r)

	return result
}
