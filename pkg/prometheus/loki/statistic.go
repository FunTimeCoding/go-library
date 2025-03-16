package loki

func (c *Client) Statistic(query string) string {
	return c.basic.Statistic(query)
}
