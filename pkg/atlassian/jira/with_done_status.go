package jira

func WithDoneStatus(v []string) OptionFunc {
	return func(c *Client) {
		c.doneStatus = v
	}
}
