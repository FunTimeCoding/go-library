package commit

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"strings"
)

func (c *Commit) formatTitle(f *option.Format) string {
	result := strings.TrimSpace(c.Title)

	if f.UseColor {
		return console.Cyan("%s", result)
	}

	return result
}
