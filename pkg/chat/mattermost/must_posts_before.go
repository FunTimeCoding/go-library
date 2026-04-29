package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/post"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (c *Client) MustPostsBefore(
	h *model.Channel,
	t time.Time,
	keep int,
) []*post.Post {
	result, e := c.PostsBefore(h, t, keep)
	errors.PanicOnError(e)

	return result
}
