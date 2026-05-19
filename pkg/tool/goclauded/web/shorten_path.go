package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func shortenPath(path string) string {
	parts := strings.Split(path, separator.Slash)

	if len(parts) <= 2 {
		return path
	}

	return strings.Join(parts[len(parts)-2:], separator.Slash)
}
