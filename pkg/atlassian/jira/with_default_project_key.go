package jira

func WithDefaultProjectKey(v string) Option {
	return func(c *Client) {
		c.defaultProjectKey = v
	}
}
