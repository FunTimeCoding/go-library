package issue

import (
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func (i *Issue) AgeWeekDays() int {
	return timeLibrary.WeekDaysSince(i.ChangeTime(), time.Now())
}
