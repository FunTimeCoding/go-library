package kestra

func WithToken(s string) Option {
	return func(c *Client) {
		c.token = s
	}
}
