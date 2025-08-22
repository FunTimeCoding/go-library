package ollama

func WithPort(s int) OptionFunc {
	return func(c *Client) {
		c.port = s
	}
}
