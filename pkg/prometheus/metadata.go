package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) Metadata(metric string) map[string][]prometheus.Metadata {
	result, e := c.client.Metadata(c.context, metric, "")
	errors.PanicOnError(e)

	return result
}
