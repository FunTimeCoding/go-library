package jira

func WithClosedStatus(v []string) Option {
	return func(c *Client) {
		c.closedStatus = v
	}
}
