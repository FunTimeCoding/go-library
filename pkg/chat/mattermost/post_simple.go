package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) PostSimple(
	h *model.Channel,
	text string,
) *model.Post {
	return c.Post(&model.Post{ChannelId: h.Id, Message: text})
}
