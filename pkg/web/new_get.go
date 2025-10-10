package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func NewGet(
	format string,
	a ...any,
) *http.Request {
	if len(a) > 0 {
		format = fmt.Sprintf(format, a...)
	}

	result, e := http.NewRequest(constant.GetMethod, format, nil)
	errors.PanicOnError(e)

	return result
}
