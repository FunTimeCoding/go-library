package chromium

func (c *Client) Close() {
	c.cancel()
	c.allocatorCancel()
}
