package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) MustAllIssues() []*issue.Issue {
	result, e := c.AllIssues()
	errors.PanicOnError(e)

	return result
}
