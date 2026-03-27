package goanalyze

import "os"

func parseFlags() (bool, bool, []string) {
	var fix, diff bool
	var patterns []string

	for _, a := range os.Args[1:] {
		switch a {
		case "--fix":
			fix = true
		case "--diff":
			diff = true
		default:
			patterns = append(patterns, a)
		}
	}

	return fix, diff, patterns
}
