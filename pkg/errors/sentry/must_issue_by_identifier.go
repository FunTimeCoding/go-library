package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustIssueByIdentifier(
	organization string,
	identifier string,
) *response.Issue {
	result, e := c.IssueByIdentifier(organization, identifier)
	errors.PanicOnError(e)

	return result
}
