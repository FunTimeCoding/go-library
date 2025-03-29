package alert

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Alert) formatTeamName(s *option.Format) string {
	var result string

	if a.Team != nil {
		return a.TeamKeyFallback()
	}

	result = NoTeam

	if s.UseColor {
		result = console.Red(result)
	}

	return result
}
