package commit

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"strings"
)

func (c *Commit) formatTitle(s *format.Settings) string {
	if s.UseColor {
		return console.Cyan("%s", strings.TrimSpace(c.Title))
	}

	return c.Title
}
