package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Posts(h *model.Channel) *model.PostList {
	result, r, e := c.client.GetPostsForChannel(
		h.Id,
		0,
		100,
		constant.EmptyEntityTag,
		true,
	)
	panicOnError(e, r)

	return result
}
