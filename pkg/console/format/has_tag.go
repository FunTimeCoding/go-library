package format

import "slices"

func (s *Settings) HasTag(v string) bool {
	if slices.Contains(s.Tags, v) {
		return true
	}

	return false
}
