package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reply(
	h *model.Channel,
	m *model.Post,
	text string,
) *model.Post {
	return c.Post(&model.Post{ChannelId: h.Id, RootId: m.Id, Message: text})
}
