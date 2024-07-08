package check_entry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (e *Entry) Format(timestamp bool) string {
	if timestamp {
		return fmt.Sprintf(
			"%s %s %s",
			e.Time.Format(time.DateMinute),
			e.Level,
			e.Text,
		)
	}

	return fmt.Sprintf("%s %s", e.Level, e.Text)
}
