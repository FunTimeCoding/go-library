package time

import (
	"fmt"
	"time"
)

func Format(t time.Time) string {
	return fmt.Sprintf("%s", t.Format(DateMinute))
}
