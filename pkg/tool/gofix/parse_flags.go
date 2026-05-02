package gofix

import "os"

func parseFlags() flags {
	var result flags

	for _, a := range os.Args[1:] {
		switch a {
		case "--diff":
			result.diff = true
		case "--survey":
			result.survey = true
		case "--rename":
			result.rename = true
		case "--summary":
			result.summary = true
		default:
			result.patterns = append(result.patterns, a)
		}
	}

	return result
}
