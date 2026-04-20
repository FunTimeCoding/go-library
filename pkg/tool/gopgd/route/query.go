package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) Query(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.QueryRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	rows, e := h.store.Query(r.Context(), body.Instance, body.Sql)
	errors.PanicOnError(e)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			server.QueryResult{Rows: toRows(rows)},
		),
	)
}
