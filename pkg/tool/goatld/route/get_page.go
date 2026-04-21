package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetPage(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePageDetail(h.confluence.Page(identifier)),
	)
}
