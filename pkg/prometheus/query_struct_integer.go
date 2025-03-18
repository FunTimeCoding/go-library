package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"github.com/funtimecoding/go-library/pkg/prometheus/result/generic/integer"
	"github.com/funtimecoding/go-library/pkg/strings"
	"time"
)

func (c *Client) QueryStructInteger(
	q string,
	fallback int,
	t time.Time,
) *integer.Result {
	result := parse.Generic(c.Query(q, t))

	if len(result) == 0 {
		return &integer.Result{Value: fallback}
	}

	return &integer.Result{
		Time:  result[0].Time,
		Value: strings.ToInteger(result[0].Value, fallback),
		Raw:   result[0],
	}
}
