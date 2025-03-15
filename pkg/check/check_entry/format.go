package check_entry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
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

	return key_value.Space(e.Level, e.Text)
}
