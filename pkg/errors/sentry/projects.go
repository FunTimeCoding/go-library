package sentry

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Projects() []sentry.Project {
	result, _, e := c.client.GetProjects()
	errors.PanicOnError(e)

	return result
}
