package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) CreateMeta(key string) *jira.CreateMetaInfo {
	result, r, e := c.client.Issue.GetCreateMeta(key)
	panicOnError(r, e)

	return result
}
