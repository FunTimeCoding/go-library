package prometheus

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Series() (v1.TSDBResult, error) {
	return c.client.TSDB(c.context)
}
