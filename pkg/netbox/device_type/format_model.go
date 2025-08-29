package device_type

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Type) formatModel(f *option.Format) string {
	result := t.Model

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
