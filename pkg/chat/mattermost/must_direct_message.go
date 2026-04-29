package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustDirectMessage(
	u *model.User,
	text string,
) *model.Post {
	result, e := c.DirectMessage(u, text)
	errors.PanicOnError(e)

	return result
}
