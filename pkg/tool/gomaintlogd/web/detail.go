package web

import "net/http"

func (s *Server) detail(
	w http.ResponseWriter,
	r *http.Request,
) {
	e := s.entryFromQuery(r)
	s.view.RenderFragment(w, detailRow(e))
}
