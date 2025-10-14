package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBefore(
	h *model.Channel,
	beforeMilli int64,
) []*post.Post {
	reference := c.FirstPostBefore(h, beforeMilli)

	if reference == nil {
		return nil
	}

	t := time.UnixMilli(beforeMilli)
	var old []*model.Post
	page := 0

	for {
		before, response, e := c.client.GetPostsBefore(
			c.context,
			h.Id,
			reference.Identifier,
			page,
			constant.PerPage,
			constant.EmptyEntityTag,
			true,
			false,
		)
		panicOnError(e, response)

		if len(before.Order) == 0 {
			break
		}

		found := false

		for _, v := range post.FromList(before, false) {
			postTime := time.UnixMilli(v.CreateAt)

			if postTime.Before(t) {
				old = append(old, v)

				if v.CreateAt < reference.Raw.CreateAt {
					reference = post.New(v)
					found = true
				}
			}
		}

		if !found {
			break
		}

		page++
	}

	result := post.NewSlice(old)
	c.Enrich(result)

	return result
}
