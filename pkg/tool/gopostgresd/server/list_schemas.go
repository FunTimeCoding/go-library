package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListSchemas(
	w http.ResponseWriter,
	r *http.Request,
	v server.ListSchemasParams,
) {
	rows, e := s.store.ListSchemas(r.Context(), v.Instance)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
