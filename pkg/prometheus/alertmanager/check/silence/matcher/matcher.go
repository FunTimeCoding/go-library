package matcher

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"regexp"
	"time"
)

func Matches(
	s *silence.Silence,
	v []*alert.Alert,
	now time.Time,
) []*alert.Alert {
	var result []*alert.Alert

	if now.Before(*s.Start) || now.After(*s.End) {
		return result
	}

	for _, a := range v {
		if matchesAlert(s, a) {
			result = append(result, a)
		}
	}

	return result
}

func matchesAlert(
	s *silence.Silence,
	a *alert.Alert,
) bool {
	for _, m := range s.Raw.Matchers {
		if !matchesLabels(m, a.Labels) {
			return false
		}
	}

	return true
}

func matchesLabels(
	m *models.Matcher,
	l models.LabelSet,
) bool {
	if false {
		fmt.Printf(
			"Matcher: name:%s value:%s regex:%v equal:%v\n",
			*m.Name,
			*m.Value,
			*m.IsRegex,
			*m.IsEqual,
		)
		fmt.Printf("LabelSet: %+v\n", l)
	}

	value, exists := l[*m.Name]

	if *m.IsRegex {
		r, e := regexp.Compile(*m.Value)

		if e != nil {
			return false
		}

		if *m.IsEqual {
			// =~
			return exists && r.MatchString(value)
		}

		// !~
		return !exists || !r.MatchString(value)
	}

	if *m.IsEqual {
		// =
		return exists && value == *m.Value
	}

	// !=
	return !exists || value != *m.Value
}
