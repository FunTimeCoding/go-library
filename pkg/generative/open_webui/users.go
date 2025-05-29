package open_webui

func (c *Client) Users() string {
	return c.basic.Request("/api/v1/users/")
}
