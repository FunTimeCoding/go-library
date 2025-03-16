package loki

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic_client/response"

func (c *Client) Query(query string) response.Data {
	return c.basic.Query(query)
}
