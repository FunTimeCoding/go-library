package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (s *Server) DeleteEntry(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int,
) {
	errors.PanicOnError(s.store.Delete(uint(identifier)))
	w.WriteHeader(http.StatusNoContent)
}
