package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/issue"
	"github.com/google/go-github/v59/github"
)

func (c *Client) Issues() []*issue.Issue {
	result, _, e := c.client.Issues.List(
		c.context,
		true,
		&github.IssueListOptions{Filter: "subscribed"},
	)
	errors.PanicOnError(e)

	return issue.NewSlice(result)
}
