package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"strings"
)

func Skipped(
	s *option.Lint,
	path string,
) bool {
	if s.Count == 0 {
		return false
	}

	for _, p := range s.Skips {
		if strings.HasPrefix(path, p) {
			return true
		}
	}

	return false
}
