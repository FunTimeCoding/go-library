package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
)

func (a *Alert) formatName(s *format.Settings) string {
	var result string
	name := a.Name

	if s.UseColor {
		name = console.Cyan(name)
	}

	if v := a.emoji(); len(v) > 0 {
		result = key_value.Space(join.Empty(v), name)
	}

	return result
}
