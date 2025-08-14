package post

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func formatLine(
	f *option.Format,
	s string,
) string {
	if s == "" {
		s = "(no text, image-only)"
	}

	if f.UseColor {
		s = console.Cyan("%s", s)
	}

	return s
}
