package verbose_transport

import "net/http"

type Transport struct {
	transport http.RoundTripper
}
