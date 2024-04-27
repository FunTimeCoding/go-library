package web_response

import "net/http"

type Response struct {
	Response     *http.Response
	BodyBytes    int
	Milliseconds int64
}
