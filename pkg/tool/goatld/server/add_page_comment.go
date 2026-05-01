package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
	"net/http"
)

func (s *Server) AddPageComment(
	w http.ResponseWriter,
	e *http.Request,
	identifier string,
) {
	var q server.CommentRequest
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&q))
	s.confluence.MustAddComment(identifier, q.Body)
	w.WriteHeader(http.StatusNoContent)
}
