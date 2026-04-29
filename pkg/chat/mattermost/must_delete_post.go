package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) MustDeletePost(p *model.Post) {
	errors.PanicOnError(c.DeletePost(p))
}
