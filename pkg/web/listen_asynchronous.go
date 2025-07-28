package web

import "net/http"

func ListenAsynchronous(m *http.ServeMux) *http.Server {
	s := &http.Server{Addr: ListenAddress, Handler: m}
	ServeAsynchronous(s)

	return s
}
