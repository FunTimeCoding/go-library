package open_webui

func (c *Client) Files() string {
	return c.basic.Get("/api/v1/files/")
}
