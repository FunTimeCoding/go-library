package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/helper"
	"time"
)

func (c *Client) LabelValues(
	label string,
	matches []string,
) []string {
	result, w, e := c.client.LabelValues(
		c.context,
		label,
		matches,
		time.Now(),
		time.Now(),
	)
	errors.PanicOnError(e)
	helper.PanicOnWarning(w)

	return helper.LabelValuesToStrings(result)
}
