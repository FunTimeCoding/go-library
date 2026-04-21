package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Project(key string) *jira.Project {
	result, r, e := c.client.Project.Get(key)
	panicOnError(r, e)

	return result
}
