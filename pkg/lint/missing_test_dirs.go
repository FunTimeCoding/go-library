package lint

import (
	"path/filepath"
	"sort"
	"strings"
)

func missingTestDirs(paths []string) []string {
	dirs := make(map[string]bool)

	for _, p := range paths {
		dirs[filepath.Dir(p)] = false
	}

	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			dirs[filepath.Dir(p)] = true
		}
	}

	var result []string

	for dir, hasTest := range dirs {
		if !hasTest {
			result = append(result, dir)
		}
	}

	sort.Strings(result)

	return result
}
