package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/query_result"
	"time"
)

func (c *Client) MustQuery(
	q string,
	t time.Time,
) *query_result.Result {
	result, e := c.Query(q, t)
	errors.PanicOnError(e)

	return result
}
