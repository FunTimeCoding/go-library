package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustSearchEvents(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) []response.Event {
	result, e := c.SearchEvents(organization, query, project, limit, cursor)
	errors.PanicOnError(e)

	return result
}
