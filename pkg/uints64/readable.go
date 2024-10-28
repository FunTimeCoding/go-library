package uints64

import "fmt"

const unit = 1024

// UnitLetter
// K: Kilo
// M: Mega
// G: Giga
// T: Tera
// P: Peta
// E: Exa
// noinspection SpellCheckingInspection
const UnitLetter = "KMGTPE"

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
