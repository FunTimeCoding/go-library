package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"github.com/funtimecoding/go-library/pkg/strings"
	"time"
)

func (c *Client) QueryFloats(
	q string,
	t time.Time,
) map[string]float64 {
	result := make(map[string]float64)

	for _, r := range parse.Generic(c.Query(q, t)) {
		result[r.Metric] = strings.ToFloatStrict(r.Value)
	}

	return result
}
