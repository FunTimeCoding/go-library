package route

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) Explain(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.ExplainRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	prefix := "EXPLAIN"

	if body.Analyze != nil && *body.Analyze {
		prefix = "EXPLAIN ANALYZE"
	}

	rows, e := h.store.Query(
		r.Context(),
		body.Instance,
		fmt.Sprintf("%s %s", prefix, body.Sql),
	)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
