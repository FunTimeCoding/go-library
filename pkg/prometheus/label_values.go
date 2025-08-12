package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/helper"
	"time"
)

func (c *Client) LabelValues(
	label string,
	matches []string,
	since time.Time,
) []string {
	result, w, e := c.client.LabelValues(
		c.context,
		label,
		matches,
		since,
		time.Now(),
	)
	errors.PanicOnError(e)
	helper.PanicOnWarning(w)

	return helper.LabelValuesToStrings(result)
}
