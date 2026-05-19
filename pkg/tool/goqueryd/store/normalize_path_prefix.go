package store

import (
	"fmt"
	"strings"
)

func normalizePathPrefix(prefix string) string {
	prefix = strings.TrimLeft(prefix, "/")

	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix = fmt.Sprintf("%s/", prefix)
	}

	return prefix
}
