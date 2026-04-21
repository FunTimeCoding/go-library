package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListTables(
	w http.ResponseWriter,
	r *http.Request,
	params server.ListTablesParams,
) {
	schema := "public"

	if params.Schema != nil {
		schema = *params.Schema
	}

	rows, e := h.store.ListTables(r.Context(), params.Instance, schema)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
