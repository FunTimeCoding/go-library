package ollama

func WithSecure(b bool) Option {
	return func(c *Client) {
		c.secure = b
	}
}
