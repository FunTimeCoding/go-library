package time

import (
	"fmt"
	"time"
)

func SecondsToClock(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	hour := duration / time.Hour
	duration -= hour * time.Hour
	minute := duration / time.Minute

	return fmt.Sprintf("%02d:%02d", hour, minute)
}
