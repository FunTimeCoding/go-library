package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) User(identifier string) *model.User {
	if !c.user.HasIdentifier(identifier) {
		result, r, e := c.client.GetUser(
			c.context,
			identifier,
			constant.EmptyEntityTag,
		)
		panicOnError(r, e)
		c.user.Add(result)
	}

	return c.user.ByIdentifier(identifier)
}
