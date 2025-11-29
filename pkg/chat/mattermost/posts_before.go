package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBefore(
	h *model.Channel,
	t time.Time,
) []*post.Post {
	page, r, e := c.client.GetPostsForChannel(
		c.context,
		h.Id,
		0,
		constant.PerPage,
		constant.EmptyEntityTag,
		false,
		false,
	)
	panicOnError(r, e)
	wrapped := post.NewSlice(post.FromList(page, false))
	c.Enrich(wrapped)
	var result []*post.Post

	for _, v := range wrapped {
		if v.Create.Before(t) {
			result = append(result, v)
		}
	}

	return result
}
