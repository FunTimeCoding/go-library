package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) TableSizes(
	w http.ResponseWriter,
	r *http.Request,
	v server.TableSizesParams,
) {
	schema := "public"

	if v.Schema != nil {
		schema = *v.Schema
	}

	rows, e := s.store.TableSizes(r.Context(), v.Instance, schema)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
