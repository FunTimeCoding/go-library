package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DirectMessage(
	u *model.User,
	text string,
) *model.Post {
	channel, r, e := c.client.CreateDirectChannel(c.context, c.Me().Id, u.Id)
	panicOnError(r, e)

	return c.PostSimple(channel, text)
}
