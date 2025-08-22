package ollama

func WithSecure(b bool) OptionFunc {
	return func(c *Client) {
		c.secure = b
	}
}
