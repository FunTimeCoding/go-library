package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseValidationErrors(body []byte) string {
	var fields map[string][]string

	if json.Unmarshal(body, &fields) != nil || len(fields) == 0 {
		return ""
	}

	var parts []string

	for field, messages := range fields {
		for _, message := range messages {
			if field == "__all__" {
				parts = append(parts, message)
			} else {
				parts = append(parts, fmt.Sprintf("%s: %s", field, message))
			}
		}
	}

	return strings.Join(parts, "; ")
}
