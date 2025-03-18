package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/helper"
	prometheus "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func (c *Client) QueryRange(
	q string,
	r prometheus.Range,
) model.Value {
	results, w, e := c.client.QueryRange(c.context, q, r)
	errors.PanicOnError(e)
	helper.PanicOnWarning(w)

	return results
}
