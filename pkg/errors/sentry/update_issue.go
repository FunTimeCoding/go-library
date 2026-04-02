package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) UpdateIssue(
	organization string,
	identifier string,
	status string,
	assignedTo string,
) *response.Issue {
	type body struct {
		Status     string `json:"status,omitempty"`
		AssignedTo string `json:"assignedTo,omitempty"`
	}
	b := c.basic.PutBytes(
		fmt.Sprintf(
			"organizations/%s/issues/%s",
			organization,
			identifier,
		),
		body{Status: status, AssignedTo: assignedTo},
	)
	var result response.Issue
	errors.PanicOnError(json.Unmarshal(b, &result))

	return &result
}
