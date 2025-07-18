package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Channel(identifier string) *model.Channel {
	result, r, e := c.client.GetChannel(
		identifier,
		constant.EmptyEntityTag,
	)
	panicOnError(e, r)

	return result
}
