package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) User() *jira.User {
	result, r, e := c.client.User.GetSelf()
	panicOnError(r, e)

	return result
}
