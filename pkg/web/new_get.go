package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func NewGet(
	locator string,
	a ...any,
) *http.Request {
	if len(a) > 0 {
		locator = fmt.Sprintf(locator, a...)
	}

	result, e := http.NewRequest(GetMethod, locator, nil)
	errors.PanicOnError(e)

	return result
}
