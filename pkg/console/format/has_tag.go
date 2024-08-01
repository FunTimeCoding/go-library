package format

import "slices"

func (s *Settings) HasTag(v string) bool {
	return slices.Contains(s.Tags, v)
}
