package web_response

import "net/http"

func New(r *http.Response) *Response {
	return &Response{Response: r}
}
