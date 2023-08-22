package web

import (
	"crypto/tls"
	"net/http"
)

func Client(secure bool) *http.Client {
	c := &http.Client{}

	if !secure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return c
}
