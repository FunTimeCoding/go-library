package model_context

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
)

func formatEvent(e event.Event) string {
	timestamp := e.CreatedAt.Format("Jan 02 15:04")

	switch e.Kind {
	case constant.Send:
		target := "all"

		if e.Target != "" {
			target = e.Target
		}

		return fmt.Sprintf(
			"%s  #%d %s → %s: %s",
			timestamp,
			e.Identifier,
			e.Name,
			target,
			e.Body,
		)
	case constant.Announce:
		return fmt.Sprintf(
			"%s  #%d %s announced: %s",
			timestamp,
			e.Identifier,
			e.Name,
			e.Body,
		)
	case constant.Complete:
		return fmt.Sprintf(
			"%s  #%d %s completed %s: %s",
			timestamp,
			e.Identifier,
			e.Name,
			e.Target,
			e.Body,
		)
	case constant.Update:
		return fmt.Sprintf(
			"%s  #%d %s updated scope: %s",
			timestamp,
			e.Identifier,
			e.Name,
			e.Body,
		)
	case constant.Summarize:
		body := e.Body

		if len(body) > 200 {
			body = fmt.Sprintf("%s…", body[:200])
		}

		return fmt.Sprintf(
			"%s  #%d %s summarized: %s",
			timestamp,
			e.Identifier,
			e.Name,
			body,
		)
	case constant.Moment:
		return fmt.Sprintf(
			"%s  #%d %s captured: %s",
			timestamp,
			e.Identifier,
			e.Name,
			e.Body,
		)
	case constant.InactivityTimeout:
		return fmt.Sprintf(
			"%s  #%d %s timed out (inactive during %s)",
			timestamp,
			e.Identifier,
			e.Name,
			e.Target,
		)
	case constant.CompleteTimeout:
		return fmt.Sprintf(
			"%s  #%d %s timed out (idle after completing)",
			timestamp,
			e.Identifier,
			e.Name,
		)
	case "register":
		return fmt.Sprintf(
			"%s  #%d %s joined",
			timestamp,
			e.Identifier,
			e.Name,
		)
	case constant.Release:
		return fmt.Sprintf(
			"%s  #%d %s left",
			timestamp,
			e.Identifier,
			e.Name,
		)
	default:
		return fmt.Sprintf(
			"%s  #%d %s %s",
			timestamp,
			e.Identifier,
			e.Name,
			e.Kind,
		)
	}
}
