package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustProjects() []response.Project {
	result, e := c.Projects()
	errors.PanicOnError(e)

	return result
}
