package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (h *Router) TransitionIssue(
	w http.ResponseWriter,
	q *http.Request,
	key string,
) {
	var body generated.TransitionRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	h.jira.Transition(key, body.TransitionIdentifier)
	w.WriteHeader(http.StatusNoContent)
}
