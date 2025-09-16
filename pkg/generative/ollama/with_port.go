package ollama

func WithPort(s int) Option {
	return func(c *Client) {
		c.port = s
	}
}
