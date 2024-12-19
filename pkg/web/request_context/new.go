package request_context

import "net/http"

func New(
	w http.ResponseWriter,
	r *http.Request,
) *Context {
	return &Context{writer: w, request: r}
}
