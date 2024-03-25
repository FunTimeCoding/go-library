package web

import (
	"bytes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func NewPost(
	locator string,
	body string,
) *http.Request {
	result, e := http.NewRequest(
		PostMethod,
		locator,
		bytes.NewBuffer([]byte(body)),
	)
	errors.PanicOnError(e)

	return result
}
