package commit

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"strings"
)

func (c *Commit) formatTitle(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", strings.TrimSpace(c.Title))
	}

	return c.Title
}
