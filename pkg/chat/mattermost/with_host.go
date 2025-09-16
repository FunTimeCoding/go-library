package mattermost

func WithHost(s string) Option {
	return func(c *Client) {
		c.host = s
	}
}
