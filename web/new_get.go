package web

import (
	"github.com/funtimecoding/go-library/errors"
	"net/http"
)

func NewGet(locator string) *http.Request {
	result, e := http.NewRequest(GetMethod, locator, nil)
	errors.PanicOnError(e)

	return result
}
