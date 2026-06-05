package go_mod

import "regexp"

var deadTagPattern = regexp.MustCompile(
	`(\S+)@(\S+): reading .+ unknown revision`,
)

func isDeadTag(stderr string) bool {
	return deadTagPattern.MatchString(stderr)
}
