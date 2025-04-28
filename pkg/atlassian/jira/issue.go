package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Issue(key string) *issue.Issue {
	result, _, e := c.client.Issue.Get(key, nil)
	errors.PanicOnError(e)

	return issue.New(result, c.IssueOption())
}
