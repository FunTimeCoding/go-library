package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) sessionDeleteAction(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(s.service.DeleteSession(identifier))
	w.Header().Set("HX-Redirect", constant.SessionsPath)
	w.WriteHeader(http.StatusOK)
}
