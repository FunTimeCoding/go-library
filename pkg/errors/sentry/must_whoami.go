package sentry

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) MustWhoami() *response.User {
	result, e := c.Whoami()
	errors.PanicOnError(e)

	return result
}
