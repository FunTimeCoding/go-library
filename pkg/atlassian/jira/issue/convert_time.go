package issue

import (
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func ConvertTime(s string) time.Time {
	return library.Parse("2006-01-02T15:04:05.000-0700", s)
}
