package open_webui

func (c *Client) Folders() string {
	return c.basic.Request("/api/v1/folders/")
}
