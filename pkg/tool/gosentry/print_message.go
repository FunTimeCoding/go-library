package gosentry

import (
	"fmt"
	"github.com/atlassian/go-sentry-api/datatype"
)

func printMessage(m *datatype.Message) {
	if m == nil {
		return
	}

	if m.Formatted != nil {
		fmt.Printf("\nMessage: %s\n", *m.Formatted)
	} else if m.Message != nil {
		fmt.Printf("\nMessage: %s\n", *m.Message)
	}
}
