package module

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (m *Module) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", m.Name)
	}

	return m.Name
}
