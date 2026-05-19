package store

import "regexp"

var hyphenatedPattern = regexp.MustCompile(
	`^[\p{L}\p{N}][\p{L}\p{N}'-]*-[\p{L}\p{N}][\p{L}\p{N}'-]*$`,
)

func isHyphenatedToken(term string) bool {
	return hyphenatedPattern.MatchString(term)
}
