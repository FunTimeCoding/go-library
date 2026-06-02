package web

import (
	"net/http"
	"strconv"
)

func (s *Server) moveDown(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)

	if e != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)

		return
	}

	s.service.MoveDown(uint(identifier))
	w.WriteHeader(http.StatusNoContent)
}
