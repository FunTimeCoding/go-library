package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (h *Router) AddIssueComment(
	w http.ResponseWriter,
	q *http.Request,
	key string,
) {
	var body generated.CommentRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	h.jira.Comment(key, body.Body)
	w.WriteHeader(http.StatusNoContent)
}
