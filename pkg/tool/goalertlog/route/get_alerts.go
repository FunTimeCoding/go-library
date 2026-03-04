package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/api"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Handler) GetAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	params api.GetAlertsParams,
) {
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			toResponse(h.store.ByName(params.Name)),
		),
	)
}
