package sentry

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Projects() []response.Project {
	var result []response.Project
	errors.PanicOnError(
		json.Unmarshal(c.basic.GetBytes("projects", nil), &result),
	)

	return result
}
