package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustOrganizationProjects(organization string) []response.Project {
	result, e := c.OrganizationProjects(organization)
	errors.PanicOnError(e)

	return result
}
