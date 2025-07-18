package mattermost

func WithToken(s string) OptionFunc {
	return func(c *Client) {
		c.token = s
	}
}
