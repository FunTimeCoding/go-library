package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (r *Router) AddIssueComment(
	w http.ResponseWriter,
	e *http.Request,
	key string,
) {
	var q server.CommentRequest
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&q))
	r.jira.Comment(key, q.Body)
	w.WriteHeader(http.StatusNoContent)
}
