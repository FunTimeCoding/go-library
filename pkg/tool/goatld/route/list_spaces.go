package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListSpaces(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.ConfluenceSpaces(h.confluence.Spaces()))
}
