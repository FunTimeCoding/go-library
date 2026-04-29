package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (s *Server) DeleteEntry(
	w http.ResponseWriter,
	_ *http.Request,
	id int,
) {
	errors.PanicOnError(s.store.Delete(uint(id)))
	w.WriteHeader(http.StatusNoContent)
}
