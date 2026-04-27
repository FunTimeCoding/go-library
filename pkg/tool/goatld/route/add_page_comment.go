package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"net/http"
)

func (r *Router) AddPageComment(
	w http.ResponseWriter,
	e *http.Request,
	identifier string,
) {
	var q server.CommentRequest
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&q))
	r.confluence.AddComment(identifier, q.Body)
	w.WriteHeader(http.StatusNoContent)
}
