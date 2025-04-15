package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (c *Client) WatchedIssues() []*issue.Issue {
	return c.Search("issue in watchedIssues() ORDER BY key ASC")
}
