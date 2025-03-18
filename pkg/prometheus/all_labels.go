package prometheus

func (c *Client) AllLabels() []string {
	return c.LabelNames([]string{})
}
