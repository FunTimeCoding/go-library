package system

import "sort"

func DirectoryContent(path string) []string {
	var result []string

	for _, e := range ReadDirectory(path) {
		result = append(result, e.Name())
	}

	sort.Strings(result)

	return result
}
