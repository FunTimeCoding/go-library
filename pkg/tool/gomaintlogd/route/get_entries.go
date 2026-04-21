package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetEntries(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetEntriesParams,
) {
	f := store.NewFilter()

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
	web.EncodeNotation(w, toResponse(entries))
}
