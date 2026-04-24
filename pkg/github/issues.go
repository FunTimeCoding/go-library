package github

import (
	"github.com/funtimecoding/go-library/pkg/github/issue"
	"github.com/google/go-github/v85/github"
)

func (c *Client) Issues() []*issue.Issue {
	result, r, e := c.client.Issues.ListAllIssues(
		c.context,
		&github.ListAllIssuesOptions{Filter: "subscribed"},
	)
	panicOnError(r, e)

	return issue.NewSlice(result)
}
