package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBeforeMilli(
	h *model.Channel,
	beforeMilli int64,
) []*post.Post {
	return c.PostsBefore(h, time.UnixMilli(beforeMilli))
}
