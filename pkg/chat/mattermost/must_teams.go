package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustTeams(userIdentifier string) []*model.Team {
	result, e := c.Teams(userIdentifier)
	errors.PanicOnError(e)

	return result
}
