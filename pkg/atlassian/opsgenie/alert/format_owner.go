package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (a *Alert) formatOwner(s *option.Format) string {
	var result string

	if a.Owner != "" {
		result = a.shortenUser(a.Owner)

		if s.UseColor {
			result = console.Green("%s", result)
		}
	} else if !s.HasTag(tag.Dense) {
		result = NoOwner

		if s.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
