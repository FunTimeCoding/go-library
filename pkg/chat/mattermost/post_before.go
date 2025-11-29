package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostBefore(
	h *model.Channel,
	t time.Time,
) *post.Post {
	pageNumber := 0

	for {
		page, r, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			pageNumber,
			constant.PerPage,
			constant.EmptyEntityTag,
			true,
			false,
		)
		panicOnError(r, e)

		if len(page.Order) == 0 {
			return nil
		}

		wrapped := post.NewSlice(post.FromList(page, false))
		c.Enrich(wrapped)

		for _, v := range wrapped {
			if v.Create.Before(t) {
				return v
			}
		}

		pageNumber++
	}
}
