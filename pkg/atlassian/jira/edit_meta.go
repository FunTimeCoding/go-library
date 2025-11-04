package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) EditMeta(key string) *jira.EditMetaInfo {
	result, r, e := c.client.Issue.GetEditMeta(&jira.Issue{Key: key})
	panicOnError(r, e)

	return result
}
