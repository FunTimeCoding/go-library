package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPost(p *model.Post) *model.Post {
	result, e := c.Post(p)
	errors.PanicOnError(e)

	return result
}
