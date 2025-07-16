package open_webui

func (c *Client) Folders() string {
	return c.basic.Get("/api/v1/folders/")
}
