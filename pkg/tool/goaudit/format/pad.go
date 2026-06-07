package format

import "fmt"

func pad(
	s string,
	width int,
) string {
	if len(s) >= width {
		return s
	}

	return fmt.Sprintf("%-*s", width, s)
}
