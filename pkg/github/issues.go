package github

import (
	"github.com/funtimecoding/go-library/pkg/github/issue"
	"github.com/google/go-github/v85/github"
)

func (c *Client) Issues() ([]*issue.Issue, error) {
	result, _, e := c.client.Issues.ListAllIssues(
		c.context,
		&github.ListAllIssuesOptions{Filter: "subscribed"},
	)

	if e != nil {
		return nil, e
	}

	return issue.NewSlice(result), nil
}
