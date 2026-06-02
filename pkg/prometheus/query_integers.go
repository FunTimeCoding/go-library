package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"github.com/funtimecoding/go-library/pkg/strings"
	"time"
)

func (c *Client) QueryIntegers(
	q string,
	t time.Time,
) map[string]int {
	result := make(map[string]int)

	for _, r := range parse.Generic(c.MustQuery(q, t).Value) {
		result[r.Metric] = strings.ToIntegerStrict(r.Value)
	}

	return result
}
