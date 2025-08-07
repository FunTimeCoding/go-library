package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func NewPropfind(locator string) *http.Request {
	result, e := http.NewRequest(constant.PropfindMethod, locator, nil)
	errors.PanicOnError(e)

	return result
}
