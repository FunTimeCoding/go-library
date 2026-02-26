package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) FindPost(identifier string) *model.Post {
	result, r, e := c.client.GetPost(
		c.context,
		identifier,
		constant.EmptyEntityTag,
	)
	panicOnError(r, e)

	return result
}
