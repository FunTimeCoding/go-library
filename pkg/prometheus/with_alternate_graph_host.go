package prometheus

func WithAlternateGraphHost(host string) Option {
	return func(c *Client) {
		c.alternateGraphHost = host
	}
}
