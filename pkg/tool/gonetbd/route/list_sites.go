package route

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListSites(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, toSites(h.client.Sites()))
}
