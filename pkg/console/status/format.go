package status

import "strings"

func (s *Status) Format() string {
	result := strings.Join(s.bubbles, " | ")

	if len(s.linesByTag) > 0 {
		for _, tag := range s.settings.Tags {
			if lines, ok := s.linesByTag[tag]; ok {
				result = extendLines(result, lines)
			}
		}
	}

	if s.settings.ShowExtended && len(s.lines) > 0 {
		for _, lines := range s.linesByTag {
			result = extendLines(result, lines)
		}

		result = extendLines(result, s.lines)
	}

	return result
}
