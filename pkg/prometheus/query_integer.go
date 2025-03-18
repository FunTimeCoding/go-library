package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"github.com/funtimecoding/go-library/pkg/strings"
	"time"
)

func (c *Client) QueryInteger(
	q string,
	t time.Time,
) int {
	result := parse.Generic(c.Query(q, t))

	if len(result) == 0 {
		return 0
	}

	return strings.ToIntegerStrict(result[0].Value)
}
