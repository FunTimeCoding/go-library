package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) sessionDeleteAction(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	s.service.DeleteSession(identifier)
	w.Header().Set("HX-Redirect", constant.SessionsPath)
	w.WriteHeader(http.StatusOK)
}
