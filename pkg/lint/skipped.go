package lint

import (
	"github.com/funtimecoding/go-library/pkg/lint/skip_configuration"
	"strings"
)

func Skipped(
	s *skip_configuration.Configuration,
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
