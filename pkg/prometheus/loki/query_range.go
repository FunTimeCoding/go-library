package loki

func (c *Client) QueryRange(query string) string {
	return c.basic.QueryRange(query)
}
