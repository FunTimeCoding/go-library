package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) SearchEvents(
	organization string,
	query string,
	project string,
	limit int,
	cursor string,
) []response.Event {
	q := map[string]string{}

	if query != "" {
		q["query"] = query
	}

	if project != "" {
		q["project"] = project
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	if cursor != "" {
		q["cursor"] = cursor
	}

	var result []response.Event
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf("organizations/%s/events", organization),
				q,
			),
			&result,
		),
	)

	return result
}
