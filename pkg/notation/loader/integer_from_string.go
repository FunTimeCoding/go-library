package loader

import "regexp"

func IntegerFromString(s string) string {
	return regexp.MustCompile(`\d+`).FindAllString(s, -1)[0]
}
