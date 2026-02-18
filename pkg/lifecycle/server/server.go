package server

import "net/http"

type Server struct {
	Mux        *http.ServeMux
	http       *http.Server
	Setup      func(*http.ServeMux)
	Middleware func(http.Handler) http.Handler
	Address    string
}
