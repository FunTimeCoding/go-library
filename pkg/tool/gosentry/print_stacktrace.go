package gosentry

import (
	"fmt"
	"github.com/atlassian/go-sentry-api/datatype"
)

func printStacktrace(s *datatype.Stacktrace) {
	fmt.Println()

	for _, f := range s.Frames {
		if f.InApp != nil && !*f.InApp {
			continue
		}

		file := "<unknown>"

		if f.Filename != nil {
			file = *f.Filename
		}

		function := "<unknown>"

		if f.Function != nil {
			function = *f.Function
		}

		line := 0.0

		if f.LineNo != nil {
			line = *f.LineNo
		}

		fmt.Printf("  %s:%g in %s\n", file, line, function)
	}
}
