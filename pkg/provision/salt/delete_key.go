package salt

func (c *Client) DeleteKey(minion string) ([]string, error) {
	return c.basic.DeleteKey(minion)
}
