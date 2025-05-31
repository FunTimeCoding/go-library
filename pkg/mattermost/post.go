package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Post(
	h *model.Channel,
	text string,
) *model.Post {
	result, _, e := c.client.CreatePost(
		&model.Post{ChannelId: h.Id, Message: text},
	)
	errors.PanicOnError(e)

	return result
}
