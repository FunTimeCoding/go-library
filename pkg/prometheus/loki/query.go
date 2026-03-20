package loki

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"

func (c *Client) Query(query string) response.QueryResult {
	return c.basic.Query(query)
}
