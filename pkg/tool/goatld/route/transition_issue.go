package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (r *Router) TransitionIssue(
	w http.ResponseWriter,
	e *http.Request,
	key string,
) {
	var q server.TransitionRequest
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&q))
	r.jira.Transition(key, q.TransitionIdentifier)
	w.WriteHeader(http.StatusNoContent)
}
