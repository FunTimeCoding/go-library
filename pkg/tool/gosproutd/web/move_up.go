package web

import (
	"net/http"
	"strconv"
)

func (s *Server) moveUp(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)

	if e != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)

		return
	}

	s.store.MoveUp(uint(identifier))
	s.notifier.Notify()
	w.WriteHeader(http.StatusNoContent)
}
