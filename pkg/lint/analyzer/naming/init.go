package naming

import (
	"fmt"
	"strings"
)

func init() {
	var entries []string

	for k, v := range suggestions {
		entries = append(
			entries,
			fmt.Sprintf("%s -> %s", k, strings.Join(v, " or ")),
		)
	}

	for k := range noSuggestion {
		entries = append(entries, fmt.Sprintf("%s (avoid)", k))
	}

	Analyzer.Doc = fmt.Sprintf(
		"flags blacklisted identifier name segments: %v",
		entries,
	)
}
