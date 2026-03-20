package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) UpdateEntry(
	w http.ResponseWriter,
	r *http.Request,
	id int,
) {
	e, f := h.store.Get(uint(id))
	errors.PanicOnError(f)
	var body generated.UpdateEntryJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	e.Action = body.Action
	e.User = body.User

	if body.System != nil {
		e.System = *body.System
	}

	if body.Service != nil {
		e.Service = *body.Service
	}

	if body.Description != nil {
		e.Description = *body.Description
	}

	if body.Timestamp != nil {
		e.Timestamp = *body.Timestamp
	}

	errors.PanicOnError(h.store.Update(e))
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(toResponse([]store.Entry{*e})[0]))
}
