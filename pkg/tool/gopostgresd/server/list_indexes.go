package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) ListIndexes(
	c context.Context,
	r server.ListIndexesRequestObject,
) (server.ListIndexesResponseObject, error) {
	if fail, okay := s.resolveInstance(r.Params.Instance); !okay {
		return server.ListIndexes400JSONResponse(*fail), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.store.ListIndexes(
		c,
		r.Params.Instance,
		schema,
		r.Params.Table,
	)

	if e != nil {
		return server.ListIndexes500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.ListIndexes200JSONResponse{Rows: toRows(rows)}, nil
}
