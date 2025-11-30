package time

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"strconv"
	"time"
)

func FromUnixNanoString(nsStr string) time.Time {
	result, e := strconv.ParseInt(nsStr, 10, 64)
	errors.PanicOnError(e)

	return time.Unix(0, result)
}
