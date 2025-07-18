package job

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (j *Job) formatName(f *option.Format) string {
	result := j.Name

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
