package request_context

import "net/http"

type Context struct {
	writer  http.ResponseWriter
	request *http.Request

	bodyRead bool
	body     string
	bodyByte []byte
}
