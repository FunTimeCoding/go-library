package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/helper"
	"github.com/prometheus/common/model"
	"time"
)

func (c *Client) Query(
	q string,
	t time.Time,
) model.Value {
	results, w, e := c.client.Query(c.context, q, t)
	errors.PanicOnError(e)
	helper.PanicOnWarning(w)

	return results
}
