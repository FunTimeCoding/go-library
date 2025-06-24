package github

import "github.com/funtimecoding/go-library/pkg/github/issue"

func (c *Client) OpenUserIssues() []*issue.Issue {
	return c.SearchIssue("is:open is:issue author:%s", *c.User().Login)
}
