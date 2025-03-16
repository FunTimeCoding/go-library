package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"slices"
)

func (a *Alert) formatSeverity(s *format.Settings) string {
	result := a.Severity

	if s.UseColor {
		if slices.Contains(constant.RedSeverities, result) {
			result = console.Red(result)
		} else if slices.Contains(constant.YellowSeverities, result) {
			result = console.Yellow(result)
		}
	}

	if result == constant.NoneSeverity {
		result = "no severity"
	}

	return result
}
