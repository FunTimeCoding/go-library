package metric

import "net/http"

func MiddlewareServer(
	address string,
	m http.Handler,
) *http.Server {
	return &http.Server{Addr: address, Handler: MiddlewareHandler(m)}
}
