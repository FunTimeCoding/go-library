package system

import "strings"

func FilesMatching(
	path string,
	prefix string,
) []string {
	var result []string

	for _, file := range Files(path) {
		if strings.HasPrefix(file, prefix) {
			result = append(result, file)
		}
	}

	return result
}
