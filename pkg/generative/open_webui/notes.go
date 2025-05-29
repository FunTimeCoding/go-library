package open_webui

func (c *Client) Notes() string {
	return c.basic.Request("/api/v1/notes/")
}
