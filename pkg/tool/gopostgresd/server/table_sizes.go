package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) TableSizes(
	c context.Context,
	r server.TableSizesRequestObject,
) (server.TableSizesResponseObject, error) {
	if fail, okay := s.resolveInstance(r.Params.Instance); !okay {
		return server.TableSizes400JSONResponse(*fail), nil
	}

	schema := "public"

	if r.Params.Schema != nil {
		schema = *r.Params.Schema
	}

	rows, e := s.store.TableSizes(c, r.Params.Instance, schema)

	if e != nil {
		return server.TableSizes500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.TableSizes200JSONResponse{Rows: toRows(rows)}, nil
}
