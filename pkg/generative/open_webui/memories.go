package open_webui

func (c *Client) Memories() string {
	return c.basic.Get("/api/v1/memories/")
}
