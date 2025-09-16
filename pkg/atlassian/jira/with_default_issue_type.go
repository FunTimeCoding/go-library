package jira

func WithDefaultIssueType(v string) OptionFunc {
	return func(c *Client) {
		c.defaultIssueType = v
	}
}
