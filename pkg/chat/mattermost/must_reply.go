package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustReply(
	h *model.Channel,
	m *model.Post,
	text string,
) *model.Post {
	result, e := c.Reply(h, m, text)
	errors.PanicOnError(e)

	return result
}
