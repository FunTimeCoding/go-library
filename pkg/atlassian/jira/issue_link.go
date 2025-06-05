package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (c *Client) IssueLink(key string) string {
	return issue.BuildLink(c.locator, key)
}
