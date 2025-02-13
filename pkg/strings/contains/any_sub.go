package contains

import "strings"

func AnySub(
	substrings []string,
	valid []string,
) bool {
	for _, element := range valid {
		for _, match := range substrings {
			if strings.Contains(element, match) {
				return true
			}
		}
	}

	return false
}
