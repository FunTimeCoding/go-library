package scan

import (
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"strings"
)

func hasMustFiles(
	v *virtual_file_system.System,
	path string,
) bool {
	for _, name := range v.MustReadDirectory(path) {
		if strings.HasPrefix(name, "must_") {
			return true
		}
	}

	return false
}
