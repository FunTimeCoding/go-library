package firefox

func WithHost(host string) Option {
	return func(c *Client) {
		c.address = host
	}
}
