package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) OrganizationProjects(o sentry.Organization) []sentry.Project {
	result, _, e := c.client.GetOrgProjects(o)
	errors.PanicOnError(e)

	return result
}
