package timeline

import "fmt"

func FormatDescription(e *Entry) string {
	switch e.Kind {
	case "announce":
		return fmt.Sprintf("%s announced: %s", e.Actor, e.Subject)
	case "complete":
		if e.Detail != "" {
			return fmt.Sprintf(
				"%s completed %s: %s",
				e.Actor,
				e.Subject,
				e.Detail,
			)
		}

		return fmt.Sprintf("%s completed %s", e.Actor, e.Subject)
	case "update":
		return fmt.Sprintf("%s updated scope: %s", e.Actor, e.Subject)
	case "send":
		return fmt.Sprintf("%s → %s", e.Actor, e.Subject)
	case "moment":
		return fmt.Sprintf("%s captured: %s", e.Actor, e.Subject)
	case "summarize":
		if e.Detail != "" {
			return fmt.Sprintf("%s summarized: %s", e.Actor, e.Detail)
		}

		return fmt.Sprintf("%s summarized", e.Actor)
	case "release":
		return fmt.Sprintf("%s left", e.Actor)
	case "register":
		return fmt.Sprintf("%s joined", e.Actor)
	case "inactivity_timeout":
		return fmt.Sprintf("%s timed out (inactive)", e.Actor)
	case "complete_timeout":
		return fmt.Sprintf("%s timed out (idle after completing)", e.Actor)
	default:
		if e.Actor != "" {
			return fmt.Sprintf("%s %s: %s", e.Actor, e.Kind, e.Subject)
		}

		return fmt.Sprintf("%s: %s", e.Kind, e.Subject)
	}
}
