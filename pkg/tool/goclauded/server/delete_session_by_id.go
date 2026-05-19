package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
	"net/http"
)

func (s *Server) DeleteSessionById(
	w http.ResponseWriter,
	r *http.Request,
	identifier string,
) {
	s.claude.Delete(identifier)
	sweep.DeleteSource(identifier)
	s.service.Store.DeleteAlias(identifier)
	w.WriteHeader(http.StatusOK)
}
