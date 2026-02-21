package assistant

func (c *Client) Stop() {
	c.connection.Stop()
}
