package uints64

import "fmt"

func Readable(bytes uint64) string {
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	index := 0
	divisor := unit

	for i := bytes / unit; i >= unit; i /= unit {
		divisor *= unit
		index++
	}

	return fmt.Sprintf(
		"%.1f %cB",
		float64(bytes)/float64(divisor),
		UnitLetter[index],
	)
}
