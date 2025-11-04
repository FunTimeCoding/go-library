package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Fields() []jira.Field {
	result, r, e := c.client.Field.GetList()
	panicOnError(r, e)

	return result
}
