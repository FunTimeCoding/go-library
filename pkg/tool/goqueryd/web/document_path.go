package web

import "strings"

func documentPath(virtualPath string) string {
	return strings.TrimPrefix(virtualPath, "qmd://")
}
