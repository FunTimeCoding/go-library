package device_role

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Role) formatName(f *option.Format) string {
	result := r.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
