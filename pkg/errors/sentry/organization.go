package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Organization(slug string) *response.Organization {
	var result response.Organization
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf("organizations/%s", slug),
				nil,
			),
			&result,
		),
	)

	return &result
}
