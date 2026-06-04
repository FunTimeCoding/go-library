package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, s.service.ListSourceTypes())
}
