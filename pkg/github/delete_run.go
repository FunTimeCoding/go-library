package github

func (c *Client) DeleteRun(
	owner string,
	name string,
	run int64,
) {
	r, e := c.client.Actions.DeleteWorkflowRun(
		c.context,
		owner,
		name,
		run,
	)
	panicOnError(r, e)
}
