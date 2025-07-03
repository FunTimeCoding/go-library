package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) OrganizationTeams(o sentry.Organization) []sentry.Team {
	result, _, e := c.client.GetTeams(o)
	errors.PanicOnError(e)

	return result
}
