package mattermost

func WithChannel(s string) Option {
	return func(c *Client) {
		c.channelName = s
	}
}
