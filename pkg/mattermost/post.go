package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Post(p *model.Post) *model.Post {
	result, _, e := c.client.CreatePost(p)
	errors.PanicOnError(e)

	return result
}
