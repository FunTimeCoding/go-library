package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Project(
	organization string,
	projectSlug string,
) *response.Project {
	var result response.Project
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf(
					"projects/%s/%s",
					organization,
					projectSlug,
				),
				nil,
			),
			&result,
		),
	)

	return &result
}
