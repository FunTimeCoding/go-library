package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func Listen(m *http.ServeMux) {
	errors.PanicOnError(http.ListenAndServe(ListenAddress, m))
}
