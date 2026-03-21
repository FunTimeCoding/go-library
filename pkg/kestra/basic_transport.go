package kestra

import "net/http"

type basicTransport struct {
	username string
	password string
	r        http.RoundTripper
}

func (t *basicTransport) RoundTrip(e *http.Request) (*http.Response, error) {
	e.SetBasicAuth(t.username, t.password)

	return t.r.RoundTrip(e)
}
