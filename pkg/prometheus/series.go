package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) Series() v1.TSDBResult {
	result, e := c.client.TSDB(c.context)
	errors.PanicOnError(e)

	return result
}
