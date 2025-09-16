package jira

func WithDefaultProjectKey(v string) OptionFunc {
	return func(c *Client) {
		c.defaultProjectKey = v
	}
}
