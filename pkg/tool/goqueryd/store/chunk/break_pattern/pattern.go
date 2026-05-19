package break_pattern

import "regexp"

type Pattern struct {
	Pattern *regexp.Regexp
	Score   int
}
