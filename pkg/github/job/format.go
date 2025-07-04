package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (j *Job) Format(f *option.Format) string {
	return status.New(f).String(
		j.Name,
		j.Status,
		j.Conclusion,
		j.CreatedAt.Format(time.DateMinute),
		j.Hash,
	).RawList(j).Format()
}
