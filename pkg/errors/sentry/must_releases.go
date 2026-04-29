package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustReleases(
	organization string,
	query string,
	limit int,
) []response.Release {
	result, e := c.Releases(organization, query, limit)
	errors.PanicOnError(e)

	return result
}
