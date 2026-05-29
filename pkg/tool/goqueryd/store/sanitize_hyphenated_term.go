package store

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func sanitizeHyphenatedTerm(term string) string {
	var result []string

	for _, p := range split.Dash(term) {
		if t := sanitizeTerm(p); t != "" {
			result = append(result, t)
		}
	}

	return join.Space(result...)
}
