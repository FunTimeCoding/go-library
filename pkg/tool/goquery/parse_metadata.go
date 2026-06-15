package goquery

import "strings"

func parseMetadata(pairs []string) map[string]string {
	if len(pairs) == 0 {
		return nil
	}

	result := map[string]string{}

	for _, pair := range pairs {
		key, value, found := strings.Cut(pair, "=")

		if found {
			result[key] = value
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
