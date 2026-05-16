package loki

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/query_result"

func (c *Client) Query(query string) *query_result.Result {
	return c.basic.Query(query)
}
