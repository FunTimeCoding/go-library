package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
	"strconv"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, e := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	errors.PanicOnError(e)
	errors.PanicOnError(s.store.Delete(uint(id)))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
