package jira

func WithDefaultProjectName(v string) Option {
	return func(c *Client) {
		c.defaultProjectName = v
	}
}
