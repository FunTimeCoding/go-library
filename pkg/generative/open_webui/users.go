package open_webui

func (c *Client) Users() string {
	return c.basic.Get("/api/v1/users/")
}
