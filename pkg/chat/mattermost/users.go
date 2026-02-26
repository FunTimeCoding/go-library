package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) Users(
	page int,
	perPage int,
) []*model.User {
	result, r, e := c.client.GetUsers(
		c.context,
		page,
		perPage,
		constant.EmptyEntityTag,
	)
	panicOnError(r, e)

	return result
}
