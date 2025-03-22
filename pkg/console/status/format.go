package status

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"strings"
)

func (s *Status) Format() string {
	result := key_value.Empty(
		spaces(s.format.Indentation),
		strings.Join(s.bubbles, " | "),
	)

	if len(s.linesByTag) > 0 {
		for _, tag := range s.format.Tags {
			if lines, okay := s.linesByTag[tag]; okay {
				result = extendLines(result, lines, s.format.Indentation)
			}
		}
	}

	if s.format.ShowExtended && len(s.lines) > 0 {
		result = extendLines(result, s.lines, s.format.Indentation)
	}

	return result
}
