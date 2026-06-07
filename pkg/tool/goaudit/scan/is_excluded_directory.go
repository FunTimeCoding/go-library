package scan

import "strings"

func isExcludedDirectory(name string) bool {
	return name == "tool" ||
		name == "generated" ||
		name == "testdata" ||
		strings.HasPrefix(name, "mock_")
}
