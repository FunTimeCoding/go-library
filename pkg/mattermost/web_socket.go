package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) WebSocket() *model.WebSocketClient {
	return c.webSocket
}
