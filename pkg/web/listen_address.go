package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func ListenAddress(m *http.ServeMux, address string) {
	errors.PanicOnError(Server(m, address).ListenAndServe())
}
