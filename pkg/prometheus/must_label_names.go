package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/label_result"
	"time"
)

func (c *Client) MustLabelNames(
	matches []string,
	since time.Time,
) *label_result.Result {
	result, e := c.LabelNames(matches, since)
	errors.PanicOnError(e)

	return result
}
