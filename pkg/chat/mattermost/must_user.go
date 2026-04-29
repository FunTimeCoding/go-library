package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustUser(identifier string) *model.User {
	result, e := c.User(identifier)
	errors.PanicOnError(e)

	return result
}
