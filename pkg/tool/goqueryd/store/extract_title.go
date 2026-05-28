package store

import (
	"path/filepath"
	"regexp"
	"strings"
)

var headingPattern = regexp.MustCompile(`(?m)^##?\s+(.+)$`)

func ExtractTitle(
	content string,
	filename string,
) string {
	match := headingPattern.FindStringSubmatch(content)

	if match != nil {
		return strings.TrimSpace(match[1])
	}

	base := filepath.Base(filename)
	extension := filepath.Ext(base)

	return strings.TrimSuffix(base, extension)
}
