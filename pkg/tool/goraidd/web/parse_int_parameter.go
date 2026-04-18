package web

import (
	"net/http"
	"strconv"
)

func parseIntParameter(r *http.Request, name string, fallback int) int {
	value := r.URL.Query().Get(name)

	if value == "" {
		return fallback
	}

	result, e := strconv.Atoi(value)

	if e != nil {
		return fallback
	}

	return result
}
