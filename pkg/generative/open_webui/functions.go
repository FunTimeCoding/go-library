package open_webui

func (c *Client) Functions() string {
	return c.basic.Request("/api/v1/functions/")
}
