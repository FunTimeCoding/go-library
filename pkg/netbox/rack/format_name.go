package rack

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Rack) formatName(f *option.Format) string {
	result := r.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
