package jira

func WithDefaultIssueType(v string) Option {
	return func(c *Client) {
		c.defaultIssueType = v
	}
}
