package web

import "net/http"

// Server Timeouts are the caller's responsibility; see lifecycle.WithProtectedServer
// for slowloris protection on request/response servers.
//
// nolint:gosec
func Server(
	h http.Handler,
	address string,
) *http.Server {
	return &http.Server{Addr: address, Handler: h}
}
