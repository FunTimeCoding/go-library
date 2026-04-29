package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostDefault(text string) *model.Post {
	result, e := c.PostDefault(text)
	errors.PanicOnError(e)

	return result
}
