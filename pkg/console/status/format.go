package status

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"strings"
)

func (s *Status) Format() string {
	result := key_value.Empty(
		spaces(s.settings.Indentation),
		strings.Join(s.bubbles, " | "),
	)

	if len(s.linesByTag) > 0 {
		for _, tag := range s.settings.Tags {
			if lines, okay := s.linesByTag[tag]; okay {
				result = extendLines(result, lines, s.settings.Indentation)
			}
		}
	}

	if s.settings.ShowExtended && len(s.lines) > 0 {
		result = extendLines(result, s.lines, s.settings.Indentation)
	}

	return result
}
