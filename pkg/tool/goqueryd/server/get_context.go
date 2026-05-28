package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetContext(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, s.service.ListContexts())
}
