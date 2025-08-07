package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
)

func NewPutBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		constant.PutMethod,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
