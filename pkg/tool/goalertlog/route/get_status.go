package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/api"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Handler) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result := api.StatusResponse{
		TotalRecords: h.store.Count(),
	}
	lastPoll := h.poller.LastPoll()

	if !lastPoll.IsZero() {
		result.LastPoll = new(lastPoll.Format(DateFormat))
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
