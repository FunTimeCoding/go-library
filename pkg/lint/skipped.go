package lint

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/lint/option"
	"path/filepath"
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
		if strings.Contains(p, ".") && !strings.Contains(p, "/") {
			matched, e := filepath.Match(p, filepath.Base(path))
			errors.PanicOnError(e)

			if matched {
				return true
			}
		} else if strings.HasPrefix(path, p) || strings.Contains(
			path,
			fmt.Sprintf("/%s", p),
		) {
			return true
		}
	}

	return false
}
