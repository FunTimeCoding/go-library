package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"
	"net/http"
)

func (s *Server) TransitionIssue(
	w http.ResponseWriter,
	e *http.Request,
	key string,
) {
	var q server.TransitionRequest
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&q))
	s.jira.MustTransition(key, q.TransitionIdentifier)
	w.WriteHeader(http.StatusNoContent)
}
