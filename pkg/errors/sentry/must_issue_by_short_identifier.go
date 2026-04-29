package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/issue"
)

func (c *Client) MustIssueByShortIdentifier(
	organization string,
	identifier string,
) *issue.Issue {
	result, e := c.IssueByShortIdentifier(organization, identifier)
	errors.PanicOnError(e)

	return result
}
