package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustUpdateChannelHeader(
	h *model.Channel,
	prefix string,
	text string,
	separator string,
) *model.Channel {
	result, e := c.UpdateChannelHeader(h, prefix, text, separator)
	errors.PanicOnError(e)

	return result
}
