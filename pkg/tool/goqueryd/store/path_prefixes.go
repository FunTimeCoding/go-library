package store

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"path"
	"strings"
)

func pathPrefixes(p string) map[string]bool {
	result := map[string]bool{separator.Slash: true}
	current := separator.Slash

	for _, segment := range strings.Split(path.Clean(p), separator.Slash) {
		if segment == "" {
			continue
		}

		current = join.Empty(current, segment, separator.Slash)
		result[current] = true
	}

	return result
}
