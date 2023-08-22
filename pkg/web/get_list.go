package web

import (
	"net/http"
	"strings"
)

func GetList(
	r *http.Request,
	key string,
) []string {
	getString := r.URL.Query().Get(key)
	var result []string

	if getString != "" {
		result = strings.Split(
			getString,
			",",
		)
	}

	return result
}
