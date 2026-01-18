package system

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func StripUntilDirectory(
	s string,
	directory string,
) string {
	parts := strings.SplitN(s, directory, 2)

	if len(parts) < 2 {
		return s
	}

	return separator.Slash + directory + parts[1]
}
