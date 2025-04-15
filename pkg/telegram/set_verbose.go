package telegram

func (c *Client) SetVerbose() {
	c.Verbose = true
	c.client.Debug = true
}
