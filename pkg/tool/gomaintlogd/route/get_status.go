package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(
		w,
		server.StatusResponse{TotalEntries: h.store.Count()},
	)
}
