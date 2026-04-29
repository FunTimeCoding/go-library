package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustFindPost(identifier string) *model.Post {
	result, e := c.FindPost(identifier)
	errors.PanicOnError(e)

	return result
}
