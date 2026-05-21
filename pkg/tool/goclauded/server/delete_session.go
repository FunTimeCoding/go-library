package server

import "net/http"

func (s *Server) DeleteSession(
	w http.ResponseWriter,
	_ *http.Request,
	callsign string,
) {
	s.service.ReleaseByCallsign(callsign)
	w.WriteHeader(http.StatusOK)
}
