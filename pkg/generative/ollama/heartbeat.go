package ollama

func (c *Client) Heartbeat() error {
	return c.client.Heartbeat(c.context)
}
