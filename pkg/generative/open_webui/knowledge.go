package open_webui

func (c *Client) Knowledge() string {
	return c.basic.Request("/api/v1/knowledge/")
}
