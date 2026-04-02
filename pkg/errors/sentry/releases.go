package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) Releases(
	organization string,
	query string,
	limit int,
) []response.Release {
	q := map[string]string{}

	if query != "" {
		q["query"] = query
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	var result []response.Release
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf("organizations/%s/releases", organization),
				q,
			),
			&result,
		),
	)

	return result
}
