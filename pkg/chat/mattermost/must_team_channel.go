package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustTeamChannel(name string) *model.Channel {
	result, e := c.TeamChannel(name)
	errors.PanicOnError(e)

	return result
}
