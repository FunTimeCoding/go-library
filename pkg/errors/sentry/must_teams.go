package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustTeams(organization string) []response.Team {
	result, e := c.Teams(organization)
	errors.PanicOnError(e)

	return result
}
