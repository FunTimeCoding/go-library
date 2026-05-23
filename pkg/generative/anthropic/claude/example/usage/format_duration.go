package usage

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	if d < 0 {
		return "0h00m"
	}

	h := int(d.Hours())
	m := int(d.Minutes()) % 60

	return fmt.Sprintf("%dh%02dm", h, m)
}
