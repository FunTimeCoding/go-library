package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/helper"
	"time"
)

func (c *Client) LabelNames(
	matches []string,
	since time.Time,
) []string {
	result, w, e := c.client.LabelNames(
		c.context,
		matches,
		since,
		time.Now(),
	)
	errors.PanicOnError(e)
	helper.PanicOnWarning(w)

	return result
}
