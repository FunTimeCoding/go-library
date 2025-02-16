package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Project(
	organization sentry.Organization,
	projectSlug string,
) sentry.Project {
	result, e := c.client.GetProject(organization, projectSlug)
	errors.PanicOnError(e)

	return result
}
