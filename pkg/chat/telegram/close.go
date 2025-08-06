package telegram

func (c *Client) Close() {
	if c.database != nil {
		c.database.Close()
	}
}
