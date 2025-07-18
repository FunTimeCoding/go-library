package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func (j *Job) formatDate(f *option.Format) string {
	var format string
	t := j.Create.Local()

	if f.HasTag(tag.Dense) && t.After(timeLibrary.Midnight(time.Now())) {
		format = timeLibrary.HourMinute
	} else {
		format = timeLibrary.DateMinute
	}

	return t.Format(format)
}
