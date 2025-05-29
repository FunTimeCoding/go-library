package open_webui

func (c *Client) Files() string {
	return c.basic.Request("/api/v1/files/")
}
