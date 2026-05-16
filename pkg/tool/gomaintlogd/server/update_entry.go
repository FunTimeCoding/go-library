package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) UpdateEntry(
	w http.ResponseWriter,
	r *http.Request,
	identifier int,
) {
	e, f := s.store.Get(uint(identifier))
	errors.PanicOnError(f)
	var body server.UpdateEntryJSONRequestBody
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

	errors.PanicOnError(s.store.Update(e))
	web.EncodeNotation(w, toResponse([]entry.Entry{*e})[0])
}
