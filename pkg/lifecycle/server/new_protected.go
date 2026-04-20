package server

import "net/http"

func NewProtected(
	m *http.ServeMux,
	setup func(*http.ServeMux),
	address string,
) *Server {
	return &Server{Mux: m, Setup: setup, Address: address, protected: true}
}
