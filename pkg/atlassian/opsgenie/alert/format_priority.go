package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"slices"
)

func (a *Alert) formatPriority(f *option.Format) string {
	var result string

	if f.UseColor {
		if a.Priority == alert.P1 {
			result = console.Red(string(a.Priority))
		} else if slices.Contains(
			[]alert.Priority{alert.P2, alert.P3},
			a.Priority,
		) {
			result = console.Yellow(string(a.Priority))
		} else {
			result = string(a.Priority)
		}
	} else {
		result = string(a.Priority)
	}

	return result
}
