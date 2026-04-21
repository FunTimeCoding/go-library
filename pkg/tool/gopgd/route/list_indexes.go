package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListIndexes(
	w http.ResponseWriter,
	r *http.Request,
	params server.ListIndexesParams,
) {
	schema := "public"

	if params.Schema != nil {
		schema = *params.Schema
	}

	rows, e := h.store.ListIndexes(
		r.Context(),
		params.Instance,
		schema,
		params.Table,
	)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
