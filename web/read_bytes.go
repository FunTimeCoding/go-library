package web

import (
	"github.com/funtimecoding/go-library/errors"
	"io"
	"net/http"
)

func ReadBytes(r *http.Response) []byte {
	result, e := io.ReadAll(r.Body)
	errors.PanicOnError(e)

	return result
}
