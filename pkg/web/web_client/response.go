package web_client

import "net/http"

type Response struct {
	Response     *http.Response
	BodyBytes    int
	Milliseconds int64
}
