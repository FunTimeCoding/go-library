package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) Issue(key string) *issue.Issue {
	result, r, e := c.client.Issue.Get(key, &jira.GetQueryOptions{})
	panicOnError(r, e)

	return issue.New(result, c.IssueOption())
}
