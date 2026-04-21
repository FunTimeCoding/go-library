package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (h *Router) AddPageComment(
	w http.ResponseWriter,
	q *http.Request,
	identifier string,
) {
	var r server.CommentRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&r))
	h.confluence.AddComment(identifier, r.Body)
	w.WriteHeader(http.StatusNoContent)
}
