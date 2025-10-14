package mattermost

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) PostsBeforeManual(
	h *model.Channel,
	beforeMilli int64,
	pages int,
) []*post.Post {
	t := time.UnixMilli(beforeMilli)
	var result []*post.Post

	for page := 0; page < pages; page++ {
		fmt.Printf("Page: %d\n", page)
		list, response, e := c.client.GetPostsForChannel(
			c.context,
			h.Id,
			page,
			constant.PerPage,
			constant.EmptyEntityTag,
			false,
			false,
		)
		panicOnError(e, response)

		if len(list.Order) == 0 {
			break
		}

		wrapped := post.NewSlice(post.FromList(list, false))
		c.Enrich(wrapped)
		found := false

		for _, v := range wrapped {
			if i := time.UnixMilli(v.Raw.CreateAt); i.Before(t) {
				result = append(result, v)
				found = true
			}
		}

		if !found && page > 0 {
			break
		}
	}

	return result
}
