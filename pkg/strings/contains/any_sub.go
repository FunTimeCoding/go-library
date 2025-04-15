package contains

import "strings"

func AnySub(
	substrings []string,
	valid []string,
) bool {
	for _, e := range valid {
		for _, m := range substrings {
			if strings.Contains(e, m) {
				return true
			}
		}
	}

	return false
}
