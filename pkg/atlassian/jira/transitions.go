package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Transitions(key string) []jira.Transition {
	result, r, e := c.client.Issue.GetTransitionsWithContext(
		c.context,
		key,
	)
	panicOnError(r, e)

	return result
}
