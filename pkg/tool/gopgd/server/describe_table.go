package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DescribeTable(
	w http.ResponseWriter,
	r *http.Request,
	table string,
	v server.DescribeTableParams,
) {
	schema := "public"

	if v.Schema != nil {
		schema = *v.Schema
	}

	rows, e := s.store.DescribeTable(
		r.Context(),
		v.Instance,
		schema,
		table,
	)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
