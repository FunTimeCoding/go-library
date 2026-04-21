package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (h *Router) TransitionIssue(
	w http.ResponseWriter,
	q *http.Request,
	key string,
) {
	var r server.TransitionRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&r))
	h.jira.Transition(key, r.TransitionIdentifier)
	w.WriteHeader(http.StatusNoContent)
}
