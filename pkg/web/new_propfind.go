package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func NewPropfind(locator string) *http.Request {
	result, e := http.NewRequest(PropfindMethod, locator, nil)
	errors.PanicOnError(e)

	return result
}
