package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostEvent(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request server.EventRequest

	if json.NewDecoder(r.Body).Decode(&request) != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)

		return
	}

	e := store.NewUsageEvent()
	e.Tool = request.Tool
	e.Surface = request.Surface
	e.Actor = request.Actor
	e.Outcome = request.Outcome

	if request.DurationMs != nil {
		e.DurationMillisecond = int64(*request.DurationMs)
	}

	if request.Detail != nil {
		encoded, marshalError := json.Marshal(*request.Detail)

		if marshalError == nil {
			detail := string(encoded)
			e.Detail = &detail
		}
	}

	s.store.Create(e)
	web.EncodeNotation(w, server.EventResponse{Id: int(e.ID)})
}
