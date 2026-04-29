package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) MustTrackedIssues() []*issue.Issue {
	result, e := c.TrackedIssues()
	errors.PanicOnError(e)

	return result
}
