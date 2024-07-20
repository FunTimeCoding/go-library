package web

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
)

func ReadBytes(r *http.Response) []byte {
	return system.ReadAll(r.Body)
}
