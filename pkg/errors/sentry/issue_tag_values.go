package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) IssueTagValues(
	organization string,
	identifier string,
	tag string,
	limit int,
) []response.TagValue {
	q := map[string]string{}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	var result []response.TagValue
	errors.PanicOnError(
		json.Unmarshal(
			c.basic.GetBytes(
				fmt.Sprintf(
					"organizations/%s/issues/%s/tags/%s/values",
					organization,
					identifier,
					tag,
				),
				q,
			),
			&result,
		),
	)

	return result
}
