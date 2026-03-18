package lint

import (
	"path/filepath"
	"strings"
)

func missingTestDirectories(
	paths []string,
	generatedPaths []string,
) map[string]string {
	directories := make(map[string]bool)
	representative := make(map[string]string)

	for _, p := range paths {
		d := filepath.Dir(p)

		if _, seen := directories[d]; !seen {
			directories[d] = false
			representative[d] = p
		}
	}

	for _, p := range generatedPaths {
		d := filepath.Dir(p)

		if _, seen := directories[d]; !seen {
			directories[d] = false
			representative[d] = p
		}
	}

	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			directories[filepath.Dir(p)] = true
		}
	}

	result := make(map[string]string)

	for d, hasTest := range directories {
		if !hasTest {
			result[d] = representative[d]
		}
	}

	return result
}
