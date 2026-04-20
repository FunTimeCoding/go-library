package server

import (
	"net"
	"net/http"
)

type Server struct {
	Mux        *http.ServeMux
	http       *http.Server
	Setup      func(*http.ServeMux)
	Middleware func(http.Handler) http.Handler
	Address    string
	listener   net.Listener
	protected  bool
}
