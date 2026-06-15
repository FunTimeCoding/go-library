package service

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func sessionSlug(path string) string {
	if i := strings.LastIndex(path, separator.Slash); i >= 0 {
		return path[:i]
	}

	return path
}
