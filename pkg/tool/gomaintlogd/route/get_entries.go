package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) GetEntries(
	w http.ResponseWriter,
	_ *http.Request,
	params generated.GetEntriesParams,
) {
	f := &store.Filter{}

	if params.System != nil {
		f.System = *params.System
	}

	if params.Service != nil {
		f.Service = *params.Service
	}

	if params.User != nil {
		f.User = *params.User
	}

	if params.Since != nil {
		f.Since = *params.Since
	}

	if params.Until != nil {
		f.Until = *params.Until
	}

	if params.Limit != nil {
		f.Limit = *params.Limit
	}

	entries, e := h.store.List(f)
	errors.PanicOnError(e)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(toResponse(entries)))
}
