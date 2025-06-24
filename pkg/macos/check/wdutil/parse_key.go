package wdutil

import (
	"fmt"
	"regexp"
	"strings"
)

func parseKey(
	output string,
	key string,
) string {
	m := regexp.MustCompile(
		fmt.Sprintf(`(?m)^\s*%s\s*:\s*(.+)$`, regexp.QuoteMeta(key)),
	).FindStringSubmatch(output)

	if len(m) >= 2 {
		return strings.TrimSpace(m[1])
	}

	return ""
}
