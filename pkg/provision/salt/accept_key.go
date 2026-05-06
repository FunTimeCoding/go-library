package salt

func (c *Client) AcceptKey(minion string) ([]string, error) {
	return c.client.AcceptKey(c.context, minion)
}
