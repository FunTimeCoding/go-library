package github

import (
	"github.com/funtimecoding/go-library/pkg/github/issue"
	"github.com/google/go-github/v83/github"
)

func (c *Client) Issues() []*issue.Issue {
	result, r, e := c.client.Issues.List(
		c.context,
		true,
		&github.IssueListOptions{Filter: "subscribed"},
	)
	panicOnError(r, e)

	return issue.NewSlice(result)
}
