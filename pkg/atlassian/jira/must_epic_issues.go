package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustEpicIssues(name string) []*issue.Issue {
	result, e := c.EpicIssues(name)
	errors.PanicOnError(e)

	return result
}
