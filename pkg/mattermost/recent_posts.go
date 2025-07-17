package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/post"
	"github.com/mattermost/mattermost-server/v6/model"
	"time"
)

func (c *Client) RecentPosts(
	h *model.Channel,
	sinceMilli int64,
) []*post.Post {
	list, response, e := c.client.GetPostsSince(
		h.Id,
		sinceMilli,
		true,
	)
	panicOnError(e, response)
	t := time.UnixMilli(sinceMilli)
	var posts []*model.Post

	for _, v := range post.FromList(list) {
		if v.Message == "" {
			continue
		}

		if time.UnixMilli(v.CreateAt).Before(t) {
			continue
		}

		posts = append(posts, v)
	}

	result := post.NewSlice(posts)
	c.Enrich(result)

	return result
}
