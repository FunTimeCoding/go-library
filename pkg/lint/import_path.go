package lint

import "strings"

func ImportPath(importLine string) string {
	result := strings.TrimSpace(importLine)

	if strings.HasPrefix(result, `"`) &&
		strings.HasSuffix(result, `"`) {
		return result[1 : len(result)-1]
	}

	if parts := strings.Fields(result); len(parts) == 2 {
		path := parts[1]

		if strings.HasPrefix(path, `"`) &&
			strings.HasSuffix(path, `"`) {
			return path[1 : len(path)-1]
		}

		return path
	}

	return result
}
