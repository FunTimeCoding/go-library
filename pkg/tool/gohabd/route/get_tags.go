package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.Tags(h.habitica.Tags()))
}
