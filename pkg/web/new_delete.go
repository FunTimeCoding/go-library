package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func NewDelete(locator string) *http.Request {
	result, e := http.NewRequest(DeleteMethod, locator, nil)
	errors.PanicOnError(e)

	return result
}
