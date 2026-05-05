package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) Query(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.QueryRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	rows, e := s.store.Query(r.Context(), body.Instance, body.Sql)
	errors.PanicOnError(e)
	web.EncodeNotation(w, server.QueryResult{Rows: toRows(rows)})
}
