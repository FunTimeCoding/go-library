package open_webui

func (c *Client) Knowledge() string {
	return c.basic.Get("/api/v1/knowledge/")
}
