package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"net/http"
)

func NewPatchBytes(
	locator string,
	body io.Reader,
) *http.Request {
	result, e := http.NewRequest(
		PatchMethod,
		locator,
		body,
	)
	errors.PanicOnError(e)

	return result
}
