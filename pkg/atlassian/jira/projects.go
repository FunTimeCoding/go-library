package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Projects() *jira.ProjectList {
	result, r, e := c.client.Project.GetList()
	panicOnError(r, e)

	return result
}
