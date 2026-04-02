package gosentry

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
)

func printEventEntries(entries []response.EventEntry) {
	for _, entry := range entries {
		switch entry.Type {
		case "exception":
			printException(entry.Data.Values)
		case "message":
			msg := entry.Data.Formatted
			if msg == "" {
				msg = entry.Data.Message
			}

			if msg != "" {
				fmt.Printf("Message: %s\n", msg)
			}
		}
	}
}
