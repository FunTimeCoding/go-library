package open_webui

func (c *Client) Prompts() string {
	return c.basic.Get("/api/v1/prompts/")
}
