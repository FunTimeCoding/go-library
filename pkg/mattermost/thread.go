package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/funtimecoding/go-library/pkg/mattermost/post"
	"github.com/mattermost/mattermost-server/v6/model"
)

func (c *Client) Thread(p *model.Post) []*post.Post {
	if p.ReplyCount == 0 {
		return []*post.Post{}
	}

	list, r, e := c.client.GetPostThread(
		p.Id,
		constant.EmptyEntityTag,
		false,
	)
	panicOnError(e, r)
	var posts []*model.Post

	for i, o := range post.FromList(list) {
		if i == 0 {
			continue
		}

		posts = append(posts, o)
	}

	result := post.NewSlice(posts)
	c.Enrich(result)

	return result
}
