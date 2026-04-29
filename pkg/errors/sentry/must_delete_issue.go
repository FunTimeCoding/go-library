package sentry

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteIssue(
	organization string,
	identifier string,
) {
	errors.PanicOnError(c.DeleteIssue(organization, identifier))
}
