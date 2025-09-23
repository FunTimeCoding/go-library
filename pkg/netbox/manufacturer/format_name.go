package manufacturer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (m *Manufacturer) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", m.Name)
	}

	return m.Name
}
