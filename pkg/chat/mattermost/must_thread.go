package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustThread(p *model.Post) []*post.Post {
	result, e := c.Thread(p)
	errors.PanicOnError(e)

	return result
}
