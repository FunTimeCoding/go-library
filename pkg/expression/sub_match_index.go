package expression

import "regexp"

func SubMatchIndex(
	expression string,
	search string,
	index int,
) string {
	m := regexp.MustCompile(expression).FindStringSubmatch(search)

	if len(m) > index {
		return m[index]
	}

	return ""
}
