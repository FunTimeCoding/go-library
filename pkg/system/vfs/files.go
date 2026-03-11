package vfs

import "sort"

func (v *VFS) Files() []string {
	result := make([]string, 0, len(v.files))

	for path := range v.files {
		result = append(result, path)
	}

	sort.Strings(result)

	return result
}
