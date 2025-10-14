package verbose_transport

import "net/http"

func Client() *http.Client {
	return &http.Client{
		Transport: &Transport{transport: http.DefaultTransport},
	}
}
