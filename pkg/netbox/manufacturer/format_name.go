package manufacturer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (m *Manufacturer) formatName(f *option.Format) string {
	result := m.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
