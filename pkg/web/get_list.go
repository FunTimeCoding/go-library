package web

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"net/http"
)

func GetList(
	r *http.Request,
	key string,
) []string {
	var result []string

	if s := r.URL.Query().Get(key); s != "" {
		result = strings.SplitComma(s)
	}

	return result
}
