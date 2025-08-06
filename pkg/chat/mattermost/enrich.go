package mattermost

import "github.com/funtimecoding/go-library/pkg/chat/mattermost/post"

func (c *Client) Enrich(v []*post.Post) {
	for _, p := range v {
		if p.User != nil {
			continue
		}

		p.Enrich(c.User(p.UserIdentifier))
	}
}
