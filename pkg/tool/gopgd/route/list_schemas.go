package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) ListSchemas(
	w http.ResponseWriter,
	r *http.Request,
	params server.ListSchemasParams,
) {
	rows, e := h.store.ListSchemas(r.Context(), params.Instance)
	errors.PanicOnError(e)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			server.QueryResult{Rows: toRows(rows)},
		),
	)
}
