package server

import "net/http"

func (s *Server) DeleteAlias(
	w http.ResponseWriter,
	_ *http.Request,
	session string,
) {
	s.service.Store.DeleteAlias(session)
	w.WriteHeader(http.StatusOK)
}
