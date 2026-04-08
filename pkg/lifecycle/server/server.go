package server

import (
	"net/http"
	"time"
)

type Server struct {
	Mux          *http.ServeMux
	http         *http.Server
	Setup        func(*http.ServeMux)
	Middleware   func(http.Handler) http.Handler
	Address      string
	WriteTimeout time.Duration
}
