package lint

import (
	"path/filepath"
	"strings"
)

func missingTestDirs(paths []string) map[string]string {
	dirs := make(map[string]bool)
	representative := make(map[string]string)

	for _, p := range paths {
		dir := filepath.Dir(p)

		if _, seen := dirs[dir]; !seen {
			dirs[dir] = false
			representative[dir] = p
		}
	}

	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			dirs[filepath.Dir(p)] = true
		}
	}

	result := make(map[string]string)

	for dir, hasTest := range dirs {
		if !hasTest {
			result[dir] = representative[dir]
		}
	}

	return result
}
