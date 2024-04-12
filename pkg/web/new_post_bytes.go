package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"net/http"
)

func NewPostBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		PostMethod,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
