package server

import "net/http"

func NewProtected(
	mux *http.ServeMux,
	setup func(*http.ServeMux),
	address string,
) *Server {
	return &Server{
		Mux:       mux,
		Setup:     setup,
		Address:   address,
		protected: true,
	}
}
