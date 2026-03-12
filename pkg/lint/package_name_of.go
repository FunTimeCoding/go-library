package lint

import (
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"strings"
)

func packageNameOf(
	v *virtual_file_system.System,
	path string,
) string {
	for _, line := range strings.Split(v.Read(path), "\n") {
		if strings.HasPrefix(line, "package ") {
			return strings.TrimPrefix(line, "package ")
		}
	}

	return ""
}
