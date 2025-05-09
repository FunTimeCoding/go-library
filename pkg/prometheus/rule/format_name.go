package rule

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Rule) formatName(f *option.Format) string {
	result := r.Name

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
