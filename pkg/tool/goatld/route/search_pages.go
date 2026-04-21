package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) SearchPages(
	w http.ResponseWriter,
	_ *http.Request,
	p server.SearchPagesParams,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePages(h.confluence.Search(p.Query)),
	)
}
