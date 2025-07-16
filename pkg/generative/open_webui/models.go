package open_webui

func (c *Client) Models() string {
	return c.basic.Get("/api/models")
}
