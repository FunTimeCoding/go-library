package chromium

func (c *Client) reconnectIfNeeded() {
	if c.context.Err() == nil {
		return
	}

	if e := c.reconnect(); e != nil {
		panic(e)
	}
}
