package file_identity

import "strings"

func skip(name string) bool {
	if strings.HasSuffix(name, "_test.go") {
		return true
	}

	if name == "main.go" || name == "doc.go" {
		return true
	}

	base := strings.TrimSuffix(name, ".go")

	if base == "constant" || strings.HasPrefix(base, "constant_") {
		return true
	}

	return false
}
