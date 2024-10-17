package web

import (
	"bytes"
	"net/http"
)

func NewPatch(
	locator string,
	body string,
) *http.Request {
	return NewPatchBytes(locator, bytes.NewBuffer([]byte(body)))
}
