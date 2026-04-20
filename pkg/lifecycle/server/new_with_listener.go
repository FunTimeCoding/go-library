package server

import (
	"net"
	"net/http"
)

func NewWithListener(
	mux *http.ServeMux,
	setup func(*http.ServeMux),
	l net.Listener,
) *Server {
	return &Server{Mux: mux, Setup: setup, listener: l}
}
