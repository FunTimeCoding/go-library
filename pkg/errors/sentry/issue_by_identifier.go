package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) IssueByIdentifier(
	organization string,
	identifier string,
) *response.Issue {
	var result response.Issue
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf(
					"organizations/%s/issues/%s",
					organization,
					identifier,
				),
				nil,
			),
			&result,
		),
	)

	return &result
}
