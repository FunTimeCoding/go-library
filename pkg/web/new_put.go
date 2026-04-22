package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func NewPut(
	locator string,
	body string,
) *http.Request {
	result, e := http.NewRequest(
		constant.Put,
		locator,
		bytes.NewBuffer([]byte(body)),
	)
	errors.PanicOnError(e)

	return result
}
