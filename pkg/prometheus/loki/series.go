package loki

func (c *Client) Series(series string) string {
	return c.basic.Series(series)
}
