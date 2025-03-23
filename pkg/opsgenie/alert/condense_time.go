package alert

import (
	"fmt"
	"github.com/docker/go-units"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func condenseTime(t time.Time) string {
	var clockFormat string
	local := t.Local()

	if time.Since(t) < 24*time.Hour {
		clockFormat = local.Format(timeLibrary.HourMinute)
	} else {
		clockFormat = local.Format(timeLibrary.DateMinute)
	}

	return fmt.Sprintf(
		"%s (%s ago)",
		clockFormat,
		units.HumanDuration(time.Since(t)),
	)
}
