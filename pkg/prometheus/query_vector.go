package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"time"
)

func (c *Client) QueryVector(q string) float64 {
	return parse.VectorFloatSingle(c.Query(q, time.Now().Add(-time.Hour)))
}
