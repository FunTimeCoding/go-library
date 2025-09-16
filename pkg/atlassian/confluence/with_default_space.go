package confluence

func WithDefaultSpace(s string) Option {
	return func(c *Client) {
		c.defaultSpace = s
	}
}
