package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result := server.StatusResponse{
		TotalRecords: h.store.Count(),
	}
	lastPoll := h.poller.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(DateFormat))
	}

	web.EncodeNotation(w, result)
}
