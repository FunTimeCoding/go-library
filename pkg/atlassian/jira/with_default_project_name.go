package jira

func WithDefaultProjectName(v string) OptionFunc {
	return func(c *Client) {
		c.defaultProjectName = v
	}
}
