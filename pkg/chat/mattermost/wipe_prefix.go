package mattermost

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (c *Client) WipePrefix(
	h *model.Channel,
	prefix string,
) {
	me := c.MustMe()

	for _, post := range c.MustPosts(h).Posts {
		if strings.HasPrefix(post.Message, prefix) {
			if post.UserId != me.Id {
				continue
			}

			c.MustDeletePost(post)
		}
	}
}
