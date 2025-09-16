package confluence

func WithLabel(v []string) Option {
	return func(c *Client) {
		c.labels = v
	}
}
