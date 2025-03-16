package loki

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/response"
	"time"
)

func (c *Client) QueryRange(
	query string,
	start time.Time,
	end time.Time,
) response.Data {
	return c.basic.QueryRange(query, start, end)
}
