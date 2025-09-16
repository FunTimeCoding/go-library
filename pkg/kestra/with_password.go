package kestra

func WithPassword(s string) Option {
	return func(c *Client) {
		c.password = s
	}
}
