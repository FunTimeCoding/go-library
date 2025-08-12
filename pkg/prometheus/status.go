package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) Status() v1.RuntimeinfoResult {
	result, e := c.client.Runtimeinfo(c.context)
	errors.PanicOnError(e)

	return result
}
