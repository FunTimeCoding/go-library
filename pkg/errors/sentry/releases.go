package sentry

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"strconv"
)

func (c *Client) Releases(
	organization string,
	query string,
	limit int,
) ([]response.Release, error) {
	q := map[string]string{}

	if query != "" {
		q["query"] = query
	}

	if limit > 0 {
		q["limit"] = strconv.Itoa(limit)
	}

	b, e := c.basic.Get(
		fmt.Sprintf("organizations/%s/releases", organization),
		q,
	)

	if e != nil {
		return nil, e
	}

	var result []response.Release
	f := json.Unmarshal(b, &result)

	if f != nil {
		return nil, f
	}

	return result, nil
}
