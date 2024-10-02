package printer

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/face"
)

func (p *Printer) Print(
	f face.Formattable,
	indent int,
) {
	var title string

	if p.setting.UseCompact {
		title = f.Meta().Short
	} else {
		title = f.Meta().Title
	}

	console.Print(f, title, indent, p.setting)
}
