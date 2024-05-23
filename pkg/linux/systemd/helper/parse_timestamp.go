package helper

import (
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"time"
)

func ParseTimestamp(s string) time.Time {
	result, e := time.Parse(constant.DateTime, s)

	if e != nil {
		return time.Time{}
	}

	return result
}
