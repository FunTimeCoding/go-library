package status

import "fmt"

func prefixed(
	prefix string,
	value string,
) string {
	if prefix == "" {
		return value
	}

	return fmt.Sprintf("%s: %s", prefix, value)
}
