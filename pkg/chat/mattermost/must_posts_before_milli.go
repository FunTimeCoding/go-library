package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustPostsBeforeMilli(
	h *model.Channel,
	beforeMilli int64,
	keep int,
) []*post.Post {
	result, e := c.PostsBeforeMilli(h, beforeMilli, keep)
	errors.PanicOnError(e)

	return result
}
