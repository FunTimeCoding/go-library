package server

import "net/http"

func New(
	m *http.ServeMux,
	setup func(*http.ServeMux),
	address string,
) *Server {
	return &Server{Mux: m, Setup: setup, Address: address}
}
