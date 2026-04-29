package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustChannelUsers(h *model.Channel) []*model.User {
	result, e := c.ChannelUsers(h)
	errors.PanicOnError(e)

	return result
}
