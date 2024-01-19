package go_mod

import "regexp"

func ExtractReplaces(modContent string) string {
	m := regexp.MustCompile(
		`(?s)replace (.*)`,
	).FindStringSubmatch(modContent)

	if len(m) > 0 {
		return m[0]
	}

	return ""
}
