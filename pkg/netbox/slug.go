package netbox

import (
	"regexp"
	"strings"
)

var nonSlug = regexp.MustCompile(`[^a-z0-9-]`)

func slug(name string) string {
	result := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	return nonSlug.ReplaceAllString(result, "")
}
