package model_context

import "strings"

func splitTags(raw string) []string {
	var result []string

	for _, t := range strings.Split(raw, ",") {
		t = strings.TrimSpace(t)

		if t != "" {
			result = append(result, t)
		}
	}

	return result
}
