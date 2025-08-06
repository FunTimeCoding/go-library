package mattermost

func WithChannel(s string) OptionFunc {
	return func(c *Client) {
		c.channelName = s
	}
}
