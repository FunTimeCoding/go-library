package option

import "slices"

func (f *Format) RemoveTag(v ...string) {
	var newTags []string

	for i, tag := range f.Tags {
		if slices.Contains(v, tag) {
			continue
		}

		newTags = append(newTags, f.Tags[i])
	}

	f.Tags = newTags
}
