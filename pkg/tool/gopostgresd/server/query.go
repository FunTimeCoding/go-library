package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) Query(
	c context.Context,
	r server.QueryRequestObject,
) (server.QueryResponseObject, error) {
	rows, e := s.store.Query(c, r.Body.Instance, r.Body.Sql)

	if e != nil {
		return server.Query500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.Query200JSONResponse{Rows: toRows(rows)}, nil
}
