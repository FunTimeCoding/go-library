package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) WebSocket() *model.WebSocketClient {
	return c.webSocket
}
