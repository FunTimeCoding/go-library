package conversations

import (
	"fmt"
	"time"
)

func relativeTimestamp(timestamp string) string {
	t, e := time.Parse(time.RFC3339Nano, timestamp)

	if e != nil {
		return timestamp
	}

	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		return fmt.Sprintf("%dm ago", int(d.Minutes()))
	case d < 24*time.Hour:
		return fmt.Sprintf("%dh ago", int(d.Hours()))
	default:
		return fmt.Sprintf("%dd ago", int(d.Hours()/24))
	}
}
