package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Channel(identifier string) *model.Channel {
	result, r, e := c.client.GetChannel(
		c.context,
		identifier,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
