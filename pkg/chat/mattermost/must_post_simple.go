package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostSimple(
	h *model.Channel,
	text string,
) *model.Post {
	result, e := c.PostSimple(h, text)
	errors.PanicOnError(e)

	return result
}
