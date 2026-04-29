package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustLatestEvent(
	organization string,
	issueIdentifier string,
) *response.Event {
	result, e := c.LatestEvent(organization, issueIdentifier)
	errors.PanicOnError(e)

	return result
}
