package open_webui

func (c *Client) Notes() string {
	return c.basic.Get("/api/v1/notes/")
}
