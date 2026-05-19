package store

import "strings"

func extractQueryTerms(query string) []string {
	words := strings.Fields(strings.ToLower(query))
	var result []string

	for _, w := range words {
		trimmed := strings.Trim(w, "\"'-")

		if len(trimmed) > 1 {
			result = append(result, trimmed)
		}
	}

	return result
}
