package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) SearchPages(
	w http.ResponseWriter,
	_ *http.Request,
	params server.SearchPagesParams,
) {
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			toConfluencePages(h.confluence.Search(params.Query)),
		),
	)
}
