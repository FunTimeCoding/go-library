package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/maps"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (a *Alert) formatRemainingLabels(s *format.Settings) string {
	if len(a.Remaining) == 0 {
		return ""
	}

	var result []string

	for _, k := range maps.StringKeys(a.Remaining) {
		v := a.Remaining[k]

		if s.UseColor {
			k = console.Yellow(k)
		}

		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}

	return join.Comma(result)
}
