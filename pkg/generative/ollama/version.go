package ollama

func (c *Client) Version() (string, error) {
	return c.client.Version(c.context)
}
