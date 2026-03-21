package web

import "net/http"

func (s *Server) edit(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	renderFragment(w, editForm(e))
}
