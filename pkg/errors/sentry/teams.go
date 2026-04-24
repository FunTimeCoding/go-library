package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func (c *Client) Teams(organization string) []response.Team {
	var result []response.Team
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf("organizations/%s/teams", organization),
				nil,
			),
			&result,
		),
	)

	return result
}
