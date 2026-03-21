package web

import "net/http"

func (s *Server) detail(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	renderFragment(w, detailRow(e))
}
