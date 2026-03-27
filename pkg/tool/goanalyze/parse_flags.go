package goanalyze

import "os"

func parseFlags() flags {
	var result flags

	for _, a := range os.Args[1:] {
		switch a {
		case "--fix":
			result.fix = true
		case "--diff":
			result.diff = true
		case "--survey":
			result.survey = true
		default:
			result.patterns = append(result.patterns, a)
		}
	}

	return result
}
