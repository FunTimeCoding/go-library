package mattermost

func WithTeam(s string) OptionFunc {
	return func(c *Client) {
		c.teamName = s
	}
}
