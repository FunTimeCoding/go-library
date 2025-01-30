package format

import "slices"

func (s *Settings) RemoveTag(v ...string) {
	var newTags []string

	for i, tag := range s.Tags {
		if slices.Contains(v, tag) {
			continue
		}

		newTags = append(newTags, s.Tags[i])
	}

	s.Tags = newTags
}
