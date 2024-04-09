package strings

import "strings"

func ReplaceAllSlice(
	content string,
	replaces []string,
) string {
	for k, v := range ToMap(replaces, "=") {
		content = strings.ReplaceAll(content, k, v)
	}

	return content
}
