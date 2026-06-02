package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/label_result"
	"time"
)

func (c *Client) MustLabelValues(
	label string,
	matches []string,
	since time.Time,
) *label_result.Result {
	result, e := c.LabelValues(label, matches, since)
	errors.PanicOnError(e)

	return result
}
