package tab

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (t *Tab) formatTitle(f *option.Format) string {
	title := t.Title

	if title == "" {
		title = NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
