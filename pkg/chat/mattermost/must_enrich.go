package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustEnrich(v []*post.Post) {
	errors.PanicOnError(c.Enrich(v))
}
