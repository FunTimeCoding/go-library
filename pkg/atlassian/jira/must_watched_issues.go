package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustWatchedIssues() []*issue.Issue {
	result, e := c.WatchedIssues()
	errors.PanicOnError(e)

	return result
}
