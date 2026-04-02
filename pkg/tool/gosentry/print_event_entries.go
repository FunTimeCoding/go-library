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
			m := entry.Data.Formatted

			if m == "" {
				m = entry.Data.Message
			}

			if m != "" {
				fmt.Printf("Message: %s\n", m)
			}
		}
	}
}
