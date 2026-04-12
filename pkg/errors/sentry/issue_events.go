package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) IssueEvents(
	organization string,
	identifier string,
	query string,
	limit int,
	cursor string,
) []response.Event {
	q := map[string]string{"full": "1"}

	if query != "" {
		q["query"] = query
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
				fmt.Sprintf(
					"organizations/%s/issues/%s/events",
					organization,
					identifier,
				),
				q,
			),
			&result,
		),
	)

	return result
}
