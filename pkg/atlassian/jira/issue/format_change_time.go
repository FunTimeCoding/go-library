package issue

import "github.com/funtimecoding/go-library/pkg/time"

func (i *Issue) FormatChangeTime() string {
	return i.ChangeTime().Format(time.DateMinute)
}
