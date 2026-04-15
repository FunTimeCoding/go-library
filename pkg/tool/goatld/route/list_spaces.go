package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) ListSpaces(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(toConfluenceSpaces(h.confluence.Spaces())),
	)
}
