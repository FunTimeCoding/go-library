package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"io"
	"net/http"
)

func NewPostBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(constant.Post, locator, body)
	errors.PanicOnError(e)

	return result
}
