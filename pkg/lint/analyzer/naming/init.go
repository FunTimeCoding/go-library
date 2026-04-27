package naming

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/segment"
	"strings"
)

func init() {
	var entries []string

	for k, v := range segment.Suggestions {
		var alts []string
		alts = append(alts, v.Letters...)
		alts = append(alts, v.Words...)
		entries = append(
			entries,
			fmt.Sprintf("%s -> %s", k, strings.Join(alts, " or ")),
		)
	}

	for k := range segment.NoSuggestion {
		entries = append(entries, fmt.Sprintf("%s (avoid)", k))
	}

	Analyzer.Doc = fmt.Sprintf(
		"flags blacklisted identifier name segments: %v",
		entries,
	)
}
