package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func formatEvent(e *event.Event) string {
	switch e.Kind {
	case constant.Send:
		if e.Target != "" {
			return fmt.Sprintf("%s → %s: %s", e.Name, e.Target, e.Body)
		}

		return fmt.Sprintf("%s → all: %s", e.Name, e.Body)
	case constant.Announce:
		return fmt.Sprintf("%s announced: %s", e.Name, e.Body)
	case constant.Complete:
		return fmt.Sprintf("%s completed %s: %s", e.Name, e.Target, e.Body)
	case constant.Update:
		return fmt.Sprintf("%s updated scope: %s", e.Name, e.Body)
	case constant.InactivityTimeout:
		return fmt.Sprintf(
			"%s timed out (inactive during %s)",
			e.Name,
			e.Target,
		)
	case constant.CompleteTimeout:
		return fmt.Sprintf("%s timed out (idle after completing)", e.Name)
	case constant.Release:
		return fmt.Sprintf("%s left", e.Name)
	case "register":
		return fmt.Sprintf("%s joined", e.Name)
	default:
		return fmt.Sprintf("%s %s", e.Name, e.Kind)
	}
}
