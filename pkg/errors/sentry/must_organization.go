package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganization(slug string) *response.Organization {
	result, e := c.Organization(slug)
	errors.PanicOnError(e)

	return result
}
