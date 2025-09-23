package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func (c *Client) QueryExemplar(
	q string,
	start time.Time,
	end time.Time,
) []prometheus.ExemplarQueryResult {
	results, e := c.client.QueryExemplars(c.context, q, start, end)
	errors.PanicOnError(e)

	return results
}
