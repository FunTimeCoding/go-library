package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) MustIssuesSimple(verbose bool) []*issue.Issue {
	result, e := c.IssuesSimple(verbose)
	errors.PanicOnError(e)

	return result
}
