package expression

import "regexp"

func SubMatch(
	expression string,
	search string,
) string {
	var result string
	matches := regexp.MustCompile(expression).FindStringSubmatch(search)

	if len(matches) > 1 {
		result = matches[1]
	}

	return result
}
