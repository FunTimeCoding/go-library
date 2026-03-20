package jira

func (c *Client) Unassign(key string) {
	r, e := c.client.Issue.UpdateAssigneeWithContext(
		c.context,
		key,
		nil,
	)
	panicOnError(r, e)
}
