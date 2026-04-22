package tool

import (
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format(timeLibrary.DateMinute)
}
