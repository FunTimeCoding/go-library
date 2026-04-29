package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustChannels(
	t *model.Team,
	u *model.User,
) []*model.Channel {
	result, e := c.Channels(t, u)
	errors.PanicOnError(e)

	return result
}
