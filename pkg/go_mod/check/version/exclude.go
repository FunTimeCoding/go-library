package version

import (
	"path/filepath"
	"strings"
)

func exclude(
	path string,
	skip []string,
) bool {
	for _, s := range skip {
		if strings.Contains(filepath.Base(path), s) {
			return true
		}

		if strings.Contains(path, s) {
			return true
		}
	}

	return false
}
