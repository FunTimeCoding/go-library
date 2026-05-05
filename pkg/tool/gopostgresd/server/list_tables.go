package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTables(
	w http.ResponseWriter,
	r *http.Request,
	v server.ListTablesParams,
) {
	schema := "public"

	if v.Schema != nil {
		schema = *v.Schema
	}

	rows, e := s.store.ListTables(r.Context(), v.Instance, schema)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
