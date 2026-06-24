package server

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/server"
)

func (s *Server) Explain(
	c context.Context,
	r server.ExplainRequestObject,
) (server.ExplainResponseObject, error) {
	if fail, okay := s.resolveInstance(r.Body.Instance); !okay {
		return server.Explain400JSONResponse(*fail), nil
	}

	prefix := "EXPLAIN"

	if r.Body.Analyze != nil && *r.Body.Analyze {
		prefix = "EXPLAIN ANALYZE"
	}

	rows, e := s.store.Query(
		c,
		r.Body.Instance,
		fmt.Sprintf("%s %s", prefix, r.Body.Sql),
	)

	if e != nil {
		return server.Explain500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.Explain200JSONResponse{Rows: toRows(rows)}, nil
}
