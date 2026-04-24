package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ScoreTask(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
	direction server.ScoreTaskParamsDirection,
) {
	web.EncodeNotation(
		w,
		convert.ScoreResult(
			h.habitica.Score(identifier, string(direction)),
		),
	)
}
