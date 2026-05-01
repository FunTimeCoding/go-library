package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustNewIssue(
	projectKey string,
	issueType string,
	summary string,
	description string,
) *jira.Issue {
	result, e := c.NewIssue(projectKey, issueType, summary, description)
	errors.PanicOnError(e)

	return result
}
