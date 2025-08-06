package mattermost

func WithHost(s string) OptionFunc {
	return func(c *Client) {
		c.host = s
	}
}
