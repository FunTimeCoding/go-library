package web

import "strings"

func shortenPaths(files string) string {
	lines := strings.Split(files, "\n")
	var result []string

	for _, line := range lines {
		result = append(result, shortenPath(line))
	}

	return strings.Join(result, ", ")
}
