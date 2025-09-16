package mattermost

func WithTeam(s string) Option {
	return func(c *Client) {
		c.teamName = s
	}
}
