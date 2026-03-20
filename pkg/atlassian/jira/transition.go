package jira

func (c *Client) Transition(
	key string,
	transitionIdentifier string,
) {
	r, e := c.client.Issue.DoTransitionWithContext(
		c.context,
		key,
		transitionIdentifier,
	)
	panicOnError(r, e)
}
