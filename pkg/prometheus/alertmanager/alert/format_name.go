package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
)

func (a *Alert) formatName(f *option.Format) string {
	result := a.Name

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	if f.HasTag(tag.Emoji) {
		if v := a.emoji(); len(v) > 0 {
			result = key_value.Space(join.Empty(v), result)
		}
	}

	return result
}
