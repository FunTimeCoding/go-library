package silence

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func (s *Silence) formatRule(f *option.Format) string {
	var result string

	if s.Rule != constant.UnknownRule {
		result = s.Rule
	} else {
		result = s.Match
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
