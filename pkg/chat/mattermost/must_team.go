package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustTeam(name string) *model.Team {
	result, e := c.Team(name)
	errors.PanicOnError(e)

	return result
}
