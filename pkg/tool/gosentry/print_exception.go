package gosentry

import (
	"fmt"
	"github.com/atlassian/go-sentry-api/datatype"
)

func printException(e *datatype.Exception) {
	if e.Values == nil {
		return
	}

	for _, v := range *e.Values {
		fmt.Println()

		if v.Type != nil {
			fmt.Printf("Exception: %s\n", *v.Type)
		}

		if v.Value != nil {
			fmt.Printf("Value:     %s\n", *v.Value)
		}

		if v.Stacktrace != nil {
			printStacktrace(v.Stacktrace)
		}
	}
}
