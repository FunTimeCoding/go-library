package silence

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (s *Silence) formatRule() string {
	if s.Rule != constant.UnknownRule {
		return s.Rule
	}

	return s.Match
}
