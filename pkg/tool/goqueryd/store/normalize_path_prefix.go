package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func normalizePathPrefix(prefix string) string {
	prefix = strings.TrimLeft(prefix, separator.Slash)

	if prefix != "" && !strings.HasSuffix(prefix, separator.Slash) {
		prefix = fmt.Sprintf("%s/", prefix)
	}

	return prefix
}
