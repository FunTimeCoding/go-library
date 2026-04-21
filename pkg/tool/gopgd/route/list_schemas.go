package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListSchemas(
	w http.ResponseWriter,
	r *http.Request,
	params server.ListSchemasParams,
) {
	rows, e := h.store.ListSchemas(r.Context(), params.Instance)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
