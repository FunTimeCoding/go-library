package kestra

func WithUser(s string) OptionFunc {
	return func(c *Client) {
		c.user = s
	}
}
