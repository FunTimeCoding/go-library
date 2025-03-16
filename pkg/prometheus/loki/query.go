package loki

func (c *Client) Query(query string) string {
	return c.basic.Query(query)
}
