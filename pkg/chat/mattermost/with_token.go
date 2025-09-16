package mattermost

func WithToken(s string) Option {
	return func(c *Client) {
		c.token = s
	}
}
