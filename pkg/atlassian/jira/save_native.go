package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) SaveNative(i *jira.Issue) *issue.Issue {
	result, r, e := c.client.Issue.UpdateWithContext(c.context, i)
	panicOnError(r, e)

	return c.enrichOne(issue.New(result, c.IssueOption()))
}
