package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) SaveNative(i *jira.Issue) *issue.Issue {
	result, _, e := c.client.Issue.UpdateWithContext(c.context, i)
	errors.PanicOnError(e)

	return c.enrichOne(issue.New(result, c.IssueOption()))
}
