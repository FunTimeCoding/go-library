package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) Targets() v1.TargetsResult {
	result, e := c.client.Targets(c.context)
	errors.PanicOnError(e)

	return result
}
