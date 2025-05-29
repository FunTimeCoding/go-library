package open_webui

func (c *Client) Prompts() string {
	return c.basic.Request("/api/v1/prompts/")
}
