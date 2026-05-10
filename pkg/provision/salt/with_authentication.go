package salt

func WithAuthentication(backend string) Option {
	return func(c *Client) {
		c.authentication = backend
	}
}
