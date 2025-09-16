package confluence

func WithDefaultPage(r string) Option {
	return func(c *Client) {
		c.defaultPage = r
	}
}
