package job

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (j *Job) Format(f *option.Format) string {
	return status.New(f).String(
		j.Name,
		j.Cluster,
		j.formatStatus(),
		j.formatCompletion(),
	).RawList(j.Raw).Format()
}
