package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) FirstPostBefore(
	h *model.Channel,
	beforeMilli int64,
) *post.Post {
	t := time.UnixMilli(beforeMilli)
	page := 0

	for {
		list, response, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			page,
			constant.PerPage,
			constant.EmptyEntityTag,
			true,
			false,
		)
		panicOnError(response, e)

		if len(list.Order) == 0 {
			return nil
		}

		wrapped := post.NewSlice(post.FromList(list, false))
		c.Enrich(wrapped)

		for _, v := range wrapped {
			if i := time.UnixMilli(v.Raw.CreateAt); i.Before(t) {
				return v
			}
		}

		page++
	}
}
