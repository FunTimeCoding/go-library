package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Me() *model.User {
	if c.meCache != nil {
		return c.meCache
	}

	result, r, e := c.client.GetMe(c.context, constant.EmptyEntityTag)
	panicOnError(e, r)
	c.meCache = result

	return c.meCache
}
