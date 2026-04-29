package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganizations() []response.Organization {
	result, e := c.Organizations()
	errors.PanicOnError(e)

	return result
}
