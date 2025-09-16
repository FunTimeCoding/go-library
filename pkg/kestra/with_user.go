package kestra

func WithUser(s string) Option {
	return func(c *Client) {
		c.user = s
	}
}
