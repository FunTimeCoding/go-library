package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, s.service.MustStatus())
}
