package verbose_transport

import (
	"crypto/tls"
	"net/http"
)

func New(secure bool) *http.Client {
	base := http.RoundTripper(nil)

	if secure {
		base = http.DefaultTransport
	} else {
		base = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return &http.Client{Transport: &Transport{base: base}}
}
