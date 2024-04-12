package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"net/http"
)

func NewPutBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		PutMethod,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
