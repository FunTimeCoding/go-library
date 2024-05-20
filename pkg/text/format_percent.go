package text

import "fmt"

func FormatPercent(
	text string,
	p float64,
) string {
	return fmt.Sprintf("%s (%.0f%%)", text, p)
}
