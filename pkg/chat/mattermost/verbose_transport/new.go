package verbose_transport

import "net/http"

func New() *http.Client {
	return &http.Client{Transport: &Transport{base: http.DefaultTransport}}
}
