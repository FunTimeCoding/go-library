package open_webui

func (c *Client) Chats() string {
	return c.basic.Get("/api/v1/chats/all")
}
