package mattermost

import (
	"strings"

	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) WipePrefixKeepLatest(
	h *model.Channel,
	prefix string,
) {
	me := c.Me()
	var candidates []*model.Post

	for _, p := range c.Posts(h).Posts {
		if strings.HasPrefix(p.Message, prefix) {
			if p.UserId != me.Id {
				continue
			}

			candidates = append(candidates, p)
		}
	}

	post.Sort(candidates)

	if count := len(candidates); len(candidates) > 1 {
		for i, p := range candidates {
			if i < count-1 {
				c.DeletePost(p)
			}
		}
	}
}
