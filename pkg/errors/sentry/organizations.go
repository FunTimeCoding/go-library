package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Organizations() []sentry.Organization {
	result, _, e := c.client.GetOrganizations()
	errors.PanicOnError(e)

	return result
}
