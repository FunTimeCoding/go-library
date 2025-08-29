package jira

func WithClosedStatus(v []string) OptionFunc {
	return func(c *Client) {
		c.closedStatus = v
	}
}
