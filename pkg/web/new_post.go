package web

import (
	"bytes"
	"net/http"
)

func NewPost(
	locator string,
	body string,
) *http.Request {
	return NewPostBytes(locator, bytes.NewBuffer([]byte(body)))
}
