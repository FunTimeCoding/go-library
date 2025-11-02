package verbose_transport

import "net/http"

type Transport struct {
	base http.RoundTripper
}
