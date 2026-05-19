package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostEmbed(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.store.Embed(s.ollama)
	errors.PanicOnError(e)
	web.EncodeNotation(w, result)
}
