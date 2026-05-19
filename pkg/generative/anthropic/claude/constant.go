package claude

import "regexp"

var (
	markupTagPattern = regexp.MustCompile(`<[^>]+>`)
	ansiPattern      = regexp.MustCompile(`\x1b\[[0-9;]*m`)
)
