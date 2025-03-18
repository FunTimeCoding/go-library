package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"github.com/funtimecoding/go-library/pkg/prometheus/result/generic/float"
	"github.com/funtimecoding/go-library/pkg/strings"
	"time"
)

func (c *Client) QueryStructFloat(
	q string,
	fallback float64,
	t time.Time,
) *float.Result {
	result := parse.Generic(c.Query(q, t))

	if len(result) == 0 {
		return &float.Result{Value: fallback}
	}

	return &float.Result{
		Time:  result[0].Time,
		Value: strings.ToFloat(result[0].Value, fallback),
		Raw:   result[0],
	}
}
