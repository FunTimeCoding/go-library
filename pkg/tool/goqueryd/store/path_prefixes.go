package store

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"path"
	"strings"
)

func pathPrefixes(p string) map[string]bool {
	result := map[string]bool{"/": true}
	current := "/"

	for _, segment := range strings.Split(path.Clean(p), "/") {
		if segment == "" {
			continue
		}

		current = join.Empty(current, segment, "/")
		result[current] = true
	}

	return result
}
