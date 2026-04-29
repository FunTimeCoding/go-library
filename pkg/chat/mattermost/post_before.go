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
) (*post.Post, error) {
	pageNumber := 0

	for {
		page, _, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			pageNumber,
			constant.PerPage,
			constant.EmptyEntityTag,
			true,
			false,
		)

		if e != nil {
			return nil, e
		}

		if len(page.Order) == 0 {
			return nil, nil
		}

		wrapped := post.NewSlice(post.FromList(page, false))
		f := c.Enrich(wrapped)

		if f != nil {
			return nil, f
		}

		for _, v := range wrapped {
			if v.Create.Before(t) {
				return v, nil
			}
		}

		pageNumber++
	}
}
