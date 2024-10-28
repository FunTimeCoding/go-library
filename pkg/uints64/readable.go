package uints64

import "fmt"

func Readable(bytes uint64) string {
	const unit = 1024

	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := unit, 0

	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf(
		"%.1f %cB",
		float64(bytes)/float64(div),
		"KMGTPE"[exp],
	)
}
