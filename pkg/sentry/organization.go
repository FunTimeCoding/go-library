package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Organization(slug string) sentry.Organization {
	result, e := c.client.GetOrganization(slug)
	errors.PanicOnError(e)

	return result
}
