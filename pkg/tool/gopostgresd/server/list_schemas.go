package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListSchemas(
	c context.Context,
	r server.ListSchemasRequestObject,
) (server.ListSchemasResponseObject, error) {
	rows, e := s.store.ListSchemas(c, r.Params.Instance)

	if e != nil {
		return server.ListSchemas500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListSchemas200JSONResponse{Rows: toRows(rows)}, nil
}
