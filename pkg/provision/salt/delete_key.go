package salt

func (c *Client) DeleteKey(minion string) ([]string, error) {
	return c.client.DeleteKey(c.context, minion)
}
