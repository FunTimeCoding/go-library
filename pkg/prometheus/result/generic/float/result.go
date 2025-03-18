package float

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/result/generic"
	"time"
)

type Result struct {
	Time  time.Time
	Value float64

	Raw *generic.Result
}
