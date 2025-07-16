package open_webui

func (c *Client) Functions() string {
	return c.basic.Get("/api/v1/functions/")
}
