package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustMe() *model.User {
	result, e := c.Me()
	errors.PanicOnError(e)

	return result
}
