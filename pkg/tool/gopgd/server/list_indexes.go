package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListIndexes(
	w http.ResponseWriter,
	r *http.Request,
	v server.ListIndexesParams,
) {
	schema := "public"

	if v.Schema != nil {
		schema = *v.Schema
	}

	rows, e := s.store.ListIndexes(
		r.Context(),
		v.Instance,
		schema,
		v.Table,
	)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
