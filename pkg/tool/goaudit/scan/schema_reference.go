package scan

import "strings"

func schemaReference(r errorHandlingResponse) string {
	if r.Content == nil {
		return ""
	}

	j, okay := r.Content["application/json"]

	if !okay {
		return ""
	}

	e := j.Schema.Reference

	if e == "" {
		return ""
	}

	parts := strings.Split(e, "/")

	return parts[len(parts)-1]
}
