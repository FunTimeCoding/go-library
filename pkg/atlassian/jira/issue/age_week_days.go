package issue

import (
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func (i *Issue) AgeWeekDays() int {
	return library.WeekDaysSince(i.ChangeTime(), time.Now())
}
