package confluence

func WithVerbose(v bool) Option {
	return func(c *Client) {
		c.verbose = v
	}
}
