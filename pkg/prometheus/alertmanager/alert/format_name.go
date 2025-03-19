package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
)

func (a *Alert) formatName(s *format.Settings) string {
	result := a.Name

	if s.UseColor {
		result = console.Cyan("%s", result)
	}

	if s.HasTag(tag.Emoji) {
		if v := a.emoji(); len(v) > 0 {
			result = key_value.Space(join.Empty(v), result)
		}
	}

	return result
}
