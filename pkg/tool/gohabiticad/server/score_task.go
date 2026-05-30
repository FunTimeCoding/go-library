package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) ScoreTask(
	_ context.Context,
	r server.ScoreTaskRequestObject,
) (server.ScoreTaskResponseObject, error) {
	return server.ScoreTask200JSONResponse(
		*convert.ScoreResult(
			s.habitica.MustScore(
				r.Identifier,
				string(r.Direction),
			),
		),
	), nil
}
