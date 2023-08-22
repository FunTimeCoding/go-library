package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func BasicGet(locator string) *http.Response {
	result, e := http.Get(locator)
	errors.PanicOnError(e)

	return result
}
