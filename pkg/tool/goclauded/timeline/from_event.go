package timeline

import (
	"fmt"
	"time"
)

func FromEvent(
	kind string,
	name string,
	body string,
	target string,
	createdAt time.Time,
	full bool,
) *Entry {
	timestamp := createdAt.UTC().Format(time.RFC3339)
	subject := body
	detail := ""

	if kind == "complete" {
		subject = target
		detail = body
	}

	if kind == "send" {
		to := "all"

		if target != "" {
			to = target
		}

		subject = fmt.Sprintf("%s: %s", to, body)
	}

	entry := &Entry{
		Timestamp: timestamp,
		Kind:      kind,
		Actor:     name,
		Subject:   subject,
		Detail:    detail,
	}

	if kind == "summarize" {
		if !full && len(body) > 200 {
			entry.Detail = fmt.Sprintf("%s…", body[:200])
		} else {
			entry.Detail = body
		}

		entry.Subject = ""
	}

	return entry
}
