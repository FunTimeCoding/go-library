package open_webui

func (c *Client) Memories() string {
	return c.basic.Request("/api/v1/memories/")
}
