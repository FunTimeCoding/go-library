package server

import "net/http"

func New(
	address string,
	setup func(*http.ServeMux),
) *Server {
	return &Server{
		Mux:     http.NewServeMux(),
		Setup:   setup,
		Address: address,
	}
}
