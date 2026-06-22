package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListTables(
	c context.Context,
	r server.ListTablesRequestObject,
) (server.ListTablesResponseObject, error) {
	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.store.ListTables(c, r.Params.Instance, schema)

	if e != nil {
		return server.ListTables500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListTables200JSONResponse{Rows: toRows(rows)}, nil
}
