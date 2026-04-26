package transport

import "net/http"

type Transport struct {
	username string
	password string
	r        http.RoundTripper
}
