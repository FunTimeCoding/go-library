package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (h *Router) AddIssueComment(
	w http.ResponseWriter,
	q *http.Request,
	key string,
) {
	var r server.CommentRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&r))
	h.jira.Comment(key, r.Body)
	w.WriteHeader(http.StatusNoContent)
}
