package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Whoami() *response.User {
	var result response.User
	errors.PanicOnError(
		json.Unmarshal(c.basic.GetBytes("user", nil), &result),
	)

	return &result
}
