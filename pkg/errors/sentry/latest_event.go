package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) LatestEvent(i sentry.Issue) sentry.Event {
	result, e := c.client.GetLatestEvent(i)
	errors.PanicOnError(e)

	return result
}
