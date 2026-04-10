package session

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Session) formatTabTitle(f *option.Format) string {
	title := s.TabTitle

	if title == "" {
		title = NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
