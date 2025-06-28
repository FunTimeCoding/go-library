package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (j *Job) Format(f *option.Format) string {
	return status.New(f).Integer(j.Identifier).String(
		j.Name,
		j.Create.Format(time.DateMinute),
		j.Status,
	).RawList(j.Raw).Format()
}
