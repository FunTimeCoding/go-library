package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetStats(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.Stats(h.habitica.UserStats()))
}
