package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	library "github.com/funtimecoding/go-library/pkg/time"
	"time"
)

func (j *Job) formatDate(f *option.Format) string {
	var format string
	t := j.Create.Local()

	if f.HasTag(tag.Dense) && t.After(library.Midnight(time.Now())) {
		format = library.HourMinute
	} else {
		format = library.DateMinute
	}

	return t.Format(format)
}
