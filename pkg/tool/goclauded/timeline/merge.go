package timeline

import "sort"

func Merge(
	entries []*Entry,
	limit int,
	offset int,
) []*Entry {
	sort.Slice(
		entries,
		func(i, j int) bool {
			return entries[i].Timestamp > entries[j].Timestamp
		},
	)

	if offset > len(entries) {
		return nil
	}

	if offset+limit > len(entries) {
		return entries[offset:]
	}

	return entries[offset : offset+limit]
}
