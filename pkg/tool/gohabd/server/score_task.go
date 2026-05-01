package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ScoreTask(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
	direction server.ScoreTaskParamsDirection,
) {
	web.EncodeNotation(
		w,
		convert.ScoreResult(
			s.habitica.MustScore(identifier, string(direction)),
		),
	)
}
