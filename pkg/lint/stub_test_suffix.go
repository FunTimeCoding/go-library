package lint

import "strings"

func stubTestSuffix(packageName string) string {
	parts := strings.Split(packageName, "_")

	return parts[len(parts)-1]
}
