package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (c *Client) ChannelUsers(h *model.Channel) []*model.User {
	var result []*model.User
	var page int

	for {
		users, r, e := c.client.GetUsersInChannel(
			c.context,
			h.Id,
			page,
			constant.PerPage,
			constant.EmptyEntityTag,
		)
		panicOnError(e, r)

		if len(users) == 0 {
			break
		}

		result = append(result, users...)
		page++
	}

	return result
}
