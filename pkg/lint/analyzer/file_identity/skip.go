package file_identity

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"strings"
)

func skip(name string) bool {
	if strings.HasSuffix(name, "_test.go") {
		return true
	}

	if name == "main.go" || name == "doc.go" {
		return true
	}

	base := strings.TrimSuffix(name, constant.GoExtension)

	if base == "constant" || strings.HasPrefix(base, "constant_") {
		return true
	}

	return false
}
