package server

import "net/http"

func (s *Server) DeleteSessionById(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	s.service.DeleteSession(identifier)
	w.WriteHeader(http.StatusOK)
}
