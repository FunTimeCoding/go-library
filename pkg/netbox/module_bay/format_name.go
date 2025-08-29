package module_bay

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Bay) formatName(f *option.Format) string {
	result := b.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
