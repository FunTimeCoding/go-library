package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustProject(
	organization string,
	projectSlug string,
) *response.Project {
	result, e := c.Project(organization, projectSlug)
	errors.PanicOnError(e)

	return result
}
