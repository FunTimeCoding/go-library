package kestra

func WithPassword(s string) OptionFunc {
	return func(c *Client) {
		c.password = s
	}
}
