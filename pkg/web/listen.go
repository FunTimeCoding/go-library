package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Listen(m *http.ServeMux) {
	errors.PanicOnError(http.ListenAndServe(constant.Listen, m))
}
