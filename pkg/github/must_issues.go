package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/github/issue"
)

func (c *Client) MustIssues() []*issue.Issue {
	result, e := c.Issues()
	errors.PanicOnError(e)

	return result
}
