package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func WriteBytes(
	w http.ResponseWriter,
	code int,
	b []byte,
) {
	w.WriteHeader(code)
	_, e := w.Write(b)
	errors.PanicOnError(e)
}
