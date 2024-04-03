package go_mod

import "regexp"

func ExtractReplaces(mod string) string {
	m := regexp.MustCompile("(?s)replace (.*)").FindStringSubmatch(mod)

	if len(m) > 0 {
		return m[0]
	}

	return ""
}
