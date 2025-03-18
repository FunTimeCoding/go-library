package integer

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/result/generic"
	"time"
)

type Result struct {
	Time  time.Time
	Value int

	Raw *generic.Result
}
