package server

import "net/http"

func New(
	mux *http.ServeMux,
	setup func(*http.ServeMux),
	address string,
) *Server {
	return &Server{
		Mux:     mux,
		Setup:   setup,
		Address: address,
	}
}
