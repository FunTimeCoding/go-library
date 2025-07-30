package mattermost

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (c *Client) WipePrefix(
	h *model.Channel,
	prefix string,
) {
	me := c.Me()

	for _, post := range c.Posts(h).Posts {
		if strings.HasPrefix(post.Message, prefix) {
			if post.UserId != me.Id {
				continue
			}

			c.DeletePost(post)
		}
	}
}
