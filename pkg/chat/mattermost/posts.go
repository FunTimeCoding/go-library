package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Posts(h *model.Channel) *model.PostList {
	result, r, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		100,
		constant.EmptyEntityTag,
		true,
		false,
	)
	panicOnError(e, r)

	return result
}
