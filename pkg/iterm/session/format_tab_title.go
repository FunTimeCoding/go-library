package session

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/iterm/constant"
)

func (s *Session) formatTabTitle(f *option.Format) string {
	title := s.TabTitle

	if title == "" {
		title = constant.NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
