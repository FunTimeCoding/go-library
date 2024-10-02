package status

import (
	"fmt"
	"strings"
)

func (s *Status) Format() string {
	result := fmt.Sprintf(
		"%s%s",
		spaces(s.settings.Indentation),
		strings.Join(s.bubbles, " | "),
	)

	if len(s.linesByTag) > 0 {
		for _, tag := range s.settings.Tags {
			if lines, ok := s.linesByTag[tag]; ok {
				result = extendLines(result, lines, s.settings.Indentation)
			}
		}
	}

	if s.settings.ShowExtended && len(s.lines) > 0 {
		for _, lines := range s.linesByTag {
			result = extendLines(result, lines, s.settings.Indentation)
		}

		result = extendLines(result, s.lines, s.settings.Indentation)
	}

	return result
}
