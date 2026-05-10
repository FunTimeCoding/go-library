package salt

func (c *Client) AcceptKey(minion string) ([]string, error) {
	return c.basic.AcceptKey(minion)
}
